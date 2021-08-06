package scnet

import (
	"io"
	"log"
	"net"
	"strconv"
	"time"

	proto "github.com/golang/protobuf/proto"
)

// TCPServer ...
type TCPServer struct {
	address  string
	Delegate TCPServerDelegate

	clients map[string]*Peer
}

// TCPServerDelegate ...
type TCPServerDelegate struct {
	ServerStarted      func(string)
	ServerStopped      func()
	ClientConnected    func(Peer)
	ClientDisconnected func(Peer)
	ClientTimeout      func(Peer) bool
}

// Start ...
func (s *TCPServer) Start(port int) {
	if s.address != "" {
		log.Print("already started")
		return
	}
	s.address = ":" + strconv.Itoa(port)

	l, err := net.Listen("tcp", s.address)
	if nil != err {
		log.Fatalf("fail to bind address to %d; err: %v", port, err)
	}
	defer l.Close()

	if s.Delegate.ServerStarted != nil {
		s.Delegate.ServerStarted(l.Addr().String())
	}

	s.clients = make(map[string]*Peer)

	tcpListener := l.(*net.TCPListener)
	defer tcpListener.Close()

	for {
		tcpListener.SetDeadline(time.Now().Add(time.Second * 10))

		conn, err := l.Accept()
		if nil != err {
			if err, ok := err.(*net.OpError); ok && err.Timeout() {
				continue
			}

			log.Printf("fail to accept; err: %v", err)
			break
		}
		defer conn.Close()

		var peer *Peer
		if p, ok := s.clients[conn.RemoteAddr().String()]; ok {
			peer = p
		} else {
			peer = &Peer{}
			s.clients[conn.RemoteAddr().String()] = peer
		}
		peer.conn = conn
		peer.Ping = time.Now()

		if s.Delegate.ClientConnected != nil {
			s.Delegate.ClientConnected(*peer)
		}

		go s.handler(peer)
	}

	s.address = ""
	if s.Delegate.ServerStopped != nil {
		s.Delegate.ServerStopped()
	}
}

func (s TCPServer) handler(peer *Peer) {
	for {
		peer.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		headLength := magicPacketLength + 4 + 2
		headData := make([]byte, headLength)
		n, err := peer.conn.Read(headData)
		if nil != err {
			if err == io.EOF {
				break
			}
			serr, ok := err.(net.Error)
			if ok && !serr.Timeout() {
				break
			}

			if s.Delegate.ClientTimeout != nil {
				if !s.Delegate.ClientTimeout(*peer) {
					break
				}
			}

			continue
		}

		if n != headLength {
			log.Print("Error: can't receive head byte")
			break
		}

		i := 0
		for i < magicPacketLength {
			if magicPacket[i] != headData[i] {
				log.Print("Error: invalid magic packet ")
				break
			}
			i++
		}

		packetLength := int(headData[2])<<24 | int(headData[3])<<16 | int(headData[4])<<8 | int(headData[5])
		headerLength := int(headData[6])<<8 | int(headData[7])

		packetData := make([]byte, packetLength)
		n, err = peer.conn.Read(packetData)
		if nil != err {
			if err == io.EOF {
				break
			}
			continue
		}

		if n != packetLength {
			log.Print("Error: invalid packet length")
			break
		}

		header := new(Header)
		headerBuf := packetData[0:headerLength]
		if err := proto.Unmarshal(headerBuf, header); err != nil {
			log.Println("Failed to parse Header: ", err)
			break
		}

		message := makeInstance(header.PacketType)
		if(message == nil) {
			println("message nil")
			break
		}
		bodyBuf := packetData[headerLength:packetLength]
		if err := proto.Unmarshal(bodyBuf, message); err != nil {
			log.Println("Failed to parse Message: ", err)
			break
		}

		peer.Ping = time.Now()

		peerCallbackMessage(peer, *header, message)
	}

	if s.Delegate.ClientDisconnected != nil {
		s.Delegate.ClientDisconnected(*peer)
	}

	delete(s.clients, peer.GetRemoteAddr())

	time.Sleep(time.Second * 3)
	defer peer.conn.Close()
}

// GetClient ...
func (s TCPServer) GetClient(address string) (*Peer, bool) {
	val, ok := s.clients[address]
	return val, ok
}

// Send ...
func (s *TCPServer) Send(peer *Peer, _header *Header, message proto.Message) {
	client := s.clients[peer.GetRemoteAddr()]
	if client == nil {
		log.Print("client not exist")
		return
	}

	header := &Header{}
	if _header != nil {
		header.ResOf = _header.Id
	}

	header.Id = 0
	header.PacketType = getPacketType(message)	

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
		return
	}
	serializedMessage, err := proto.Marshal(message)
	if err != nil {
		return
	}

	buffer = append(buffer, serializedHeader[:]...)
	buffer = append(buffer, serializedMessage[:]...)

	_, err = peer.conn.Write(buffer)
	if err != nil {
		println(err)
		return
	}
}

// RegistMessage ...
func (s *TCPServer) RegistMessage(message interface{}, pcb PeerCallback) {
	registMessageWithPeer(message, pcb)
}
