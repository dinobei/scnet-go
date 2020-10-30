package scnet

import (
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

// Status is represent connection status with server
type Status int

// Status value
const (
	Attached     = Status(0)
	Detached     = Status(1)
	Connecting   = Status(2)
	Connected    = Status(3)
	Disconnected = Status(4)
)

// TCPClient ...
type TCPClient struct {
	address  string
	Delegate TCPClientDelegate

	Peer *Peer

	isStop bool
	mutex  sync.RWMutex

	dedicatedCallbackMap map[int32]ClientCallback
	callbackKeyRef       int32

	status Status
}

// NewTCPClient ...
func NewTCPClient() *TCPClient {
	tcpClient := TCPClient{}
	tcpClient.dedicatedCallbackMap = map[int32]ClientCallback{}
	tcpClient.callbackKeyRef = 1000
	return &tcpClient
}

// TCPClientDelegate ...
type TCPClientDelegate struct {
	Attached     func()
	Detached     func()
	TimedOut     func(*TCPClient) bool
	Connected    func(*TCPClient)
	Connecting   func(*TCPClient)
	Disconnected func(*TCPClient)
}

// GetAddress ...
func (c *TCPClient) GetAddress() string {
	return c.address
}

// Attach ...
func (c *TCPClient) Attach(ip string, port int, timeout time.Duration) {
	c.address = ip + ":" + strconv.Itoa(port)

	c.status = Attached
	if c.Delegate.Attached != nil {
		c.Delegate.Attached()
	}

	c.Peer = &Peer{}
	for {
		c.mutex.RLock()
		if c.isStop {
			break
		}
		c.mutex.RUnlock()

		c.Peer.conn, _ = net.DialTimeout("tcp", c.address, timeout)

		if c.Peer.conn == nil {
			// timeout
			c.status = Connecting
			if c.Delegate.Connecting != nil {
				c.Delegate.Connecting(c)
			}

			time.Sleep(timeout)
			continue
		}

		c.status = Connected
		if c.Delegate.Connected != nil {
			c.Delegate.Connected(c)
		}

		c.MainLoop()

		c.status = Disconnected
		if c.Delegate.Disconnected != nil {
			c.Delegate.Disconnected(c)
		}

		c.Peer.conn.Close()
		c.Peer.conn = nil
	}

	c.status = Detached
	if c.Delegate.Detached != nil {
		c.Delegate.Detached()
	}
}

// Detach ...
func (c *TCPClient) Detach() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop {
		return
	}
	c.isStop = true

	if c.Peer.conn != nil {
		c.Peer.conn.Close()
	}
}

// Send ...
func (c *TCPClient) Send(header *Header, message proto.Message, callback ClientCallback) bool {
	if c.status != Connected {
		return false
	}

	packetType := getPacketType(proto.MessageName(message))
	if header == nil {
		header = &Header{}
	} else {
		if header.ReqCb > 0 {
			header.ResCb = header.ReqCb
		} else {
			header.ResCb = 0
		}
	}

	header.PacketType = packetType
	if callback != nil {
		// set callback
		header.ReqCb = c.callbackKeyRef
		c.dedicatedCallbackMap[c.callbackKeyRef] = callback
		c.callbackKeyRef++
	} else {
		header.ReqCb = 0
	}

	headerSize := proto.Size(header)
	packetSize := proto.Size(message) + headerSize
	buffer := make([]byte, magicPacketLength+4+2, magicPacketLength+4+2+packetSize)
	buffer[0] = magicPacket[0]
	buffer[1] = magicPacket[1]
	buffer[2] = byte(packetSize>>24) & 0xFF
	buffer[3] = byte(packetSize>>16) & 0xFF
	buffer[4] = byte(packetSize>>8) & 0xFF
	buffer[5] = byte(packetSize) & 0xFF
	buffer[6] = byte(headerSize>>8) & 0xFF
	buffer[7] = byte(headerSize) & 0xFF

	serializedHeader, err := proto.Marshal(header)
	if err != nil {
		return false
	}
	serializedMessage, err := proto.Marshal(message)
	if err != nil {
		return false
	}

	buffer = append(buffer, serializedHeader[:]...)
	buffer = append(buffer, serializedMessage[:]...)

	_, err = c.Peer.conn.Write(buffer)
	if err != nil {
		println(err)
		return false
	}

	return true
}

// MainLoop ...
func (c *TCPClient) MainLoop() {
	headLength := magicPacketLength + 4 + 2
	headData := make([]byte, headLength)
	for {
		// TODO: Check timeout working
		c.Peer.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		n, err := c.Peer.conn.Read(headData)
		if nil != err {
			if err == io.EOF {
				break
			}
			netErr, ok := err.(net.Error)
			if ok && !netErr.Timeout() {
				break
			}

			if c.Delegate.TimedOut != nil {
				if !c.Delegate.TimedOut(c) {
					break
				}
			}

			continue
		}

		if n != headLength {
			log.Print("Error: can't receive head byte")
			return
		}

		i := 0
		for i < magicPacketLength {
			if magicPacket[i] != headData[i] {
				log.Print("Error: invalid magic packet ")
				return
			}
			i++
		}

		packetLength := int(headData[2])<<24 | int(headData[3])<<16 | int(headData[4])<<8 | int(headData[5])
		headerLength := int(headData[6])<<8 | int(headData[7])

		// TODO: remove make() statement
		packetData := make([]byte, packetLength)
		n, err = c.Peer.conn.Read(packetData)
		if nil != err {
			if err == io.EOF {
				break
			}
			continue
		}

		c.Peer.Ping = time.Now()

		if n != packetLength {
			log.Print("Error: invalid packet length")
			time.Sleep(time.Second * 1)
			return
		}

		header := new(Header)
		headerBuf := packetData[0:headerLength]
		if err := proto.Unmarshal(headerBuf, header); err != nil {
			log.Println("Failed to parse Header: ", err)
			return
		}

		message := makeInstance(header.PacketType)
		messageBuf := packetData[headerLength:packetLength]
		if err := proto.Unmarshal(messageBuf, message); err != nil {
			log.Println("Failed to parse Header: ", err)
			return
		}

		dedicatedCallback := c.dedicatedCallbackMap[header.ResCb]
		if dedicatedCallback != nil {
			dedicatedCallback(c, *header, message)
		} else {
			clientCallbackMessage(c, *header, message)
		}
	}
}

// RegistMessage ...
func (c *TCPClient) RegistMessage(packetType uint32, message interface{}, ccb ClientCallback) {
	registMessage(packetType, message, ccb)
}
