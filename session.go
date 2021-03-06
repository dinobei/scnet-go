package scnet

import (
	"io"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
)

type readHeaderError struct{}

func (m *readHeaderError) Error() string {
	return "failed reading header"
}

type unknownProtobufPacketError struct{}

func (m *unknownProtobufPacketError) Error() string {
	return "unknown protobuf packet"
}

type notSupportedDataError struct{}

func (m *notSupportedDataError) Error() string {
	return "not supported data"
}

// Peer ...
type Peer struct {
	conn net.Conn
	Ping time.Time
}

// GetRemoteAddr ...
func (peer Peer) GetRemoteAddr() string {
	return peer.conn.RemoteAddr().String()
}

// GetLocalAddr ...
func (peer Peer) GetLocalAddr() string {
	return peer.conn.LocalAddr().String()
}

// GetFormattedPing ...
func (peer Peer) GetFormattedPing() string {
	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := peer.Ping.In(loc)
	return kst.Format("2006-01-02 15:04:05")
}

func recvHeader(conn net.Conn) (header *messageHeader, err error) {
	data := make([]byte, 2)
	n, err := conn.Read(data)
	if nil != err {
		return nil, err
	}

	if n != 2 {
		return nil, &readHeaderError{}
	}

	// magic packet validation
	i := 0
	for i < magicPacketLength {
		if magicPacket[i] != data[i] {
			return nil, &readHeaderError{}
		}
		i++
	}

	receivedHeaderComponent := 0
	headerBuffer := make([][]byte, headerElements)

	buf := make([]byte, 1)
	headerBuffer[receivedHeaderComponent] = make([]byte, 0)
	for {
		_, err := conn.Read(buf)
		if nil != err {
			return nil, err
		}

		headerBuffer[receivedHeaderComponent] = append(headerBuffer[receivedHeaderComponent], buf[0])

		if buf[0]&0xFF > 127 {
			continue
		}

		receivedHeaderComponent++
		if receivedHeaderComponent == headerElements {
			break
		}

		headerBuffer[receivedHeaderComponent] = make([]byte, 0)
	}

	header = new(messageHeader)
	header.dataSize = decodeVarint(headerBuffer[0])
	header.packetType = decodeVarint(headerBuffer[1])
	header.messageType = decodeVarint(headerBuffer[2])
	header.cryptType = decodeVarint(headerBuffer[4])
	header.connectionID = decodeVarint(headerBuffer[3])

	return header, nil
}

func recvRawBody(conn net.Conn, header messageHeader) ([]byte, error) {
	responseSize := int(header.dataSize)
	if responseSize == 0 {
		return nil, nil
	}

	buf := make([]byte, responseSize)
	tmp := make([]byte, responseSize)

	for {
		n, err := conn.Read(tmp)
		if nil != err {
			return nil, err
		}

		buf = append(buf, tmp[:n]...)
		if len(buf) >= responseSize {
			break
		}
	}

	return buf, nil
}

func recvProtobufBody(conn net.Conn, header messageHeader) (proto.Message, error) {
	msg := makeInstance(header.packetType)
	if msg == nil {
		println("unknown protobuf packet")
		return nil, &unknownProtobufPacketError{}
	}
	field := msg.(proto.Message)

	responseSize := int(header.dataSize)
	if responseSize > 0 {
		buf := make([]byte, 0, responseSize)
		tmp := make([]byte, responseSize)

		for {
			n, err := conn.Read(tmp)
			if nil != err {
				return nil, err
			}

			buf = append(buf, tmp[:n]...)
			if len(buf) >= responseSize {
				break
			}
		}

		err := proto.Unmarshal(buf, field)
		if err != nil {
			return nil, err
		}

	}

	return field, nil
}

func makeHeader(messageType int32, packetType int32, dataSize int) []byte {
	size := magicPacketLength + maxPacketHeaderSize + dataSize
	buffer := make([]byte, magicPacketLength, size)

	for i := 0; i < magicPacketLength; i++ {
		buffer[i] = magicPacket[i]
	}

	buffer = append(buffer, encodeVarint(int32(dataSize))[:]...) // data size
	buffer = append(buffer, encodeVarint(packetType)[:]...)      // packet type
	buffer = append(buffer, encodeVarint(messageType)[:]...)     // message type
	buffer = append(buffer, encodeVarint(0)[:]...)               // crypt type
	buffer = append(buffer, encodeVarint(0)[:]...)               // connectionID

	return buffer
}

// RawbyteData ...
type RawbyteData struct {
	PacketType int32
	Buffer     []byte
}

// Send ...
func Send(peer Peer, data interface{}) error {
	if casted, ok := data.(proto.Message); ok {
		return sendProtobuf(peer, casted)
	} else if casted, ok := data.(RawbyteData); ok {
		return sendRawByte(peer, casted.PacketType, casted.Buffer)
	}

	return &notSupportedDataError{}

}

func sendProtobuf(peer Peer, message proto.Message) error {
	packetType := getPacketType(proto.MessageName(message))

	buffer := makeHeader(mtProtobuf, packetType, proto.Size(message))

	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	buffer = append(buffer, data[:]...)

	_, err = peer.conn.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}

func sendRawByte(peer Peer, packetType int32, data []byte) error {
	buffer := makeHeader(mtRawbyte, packetType, len(data))
	buffer = append(buffer, data[:]...)

	_, err := peer.conn.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}

func connHandler(peer *Peer, checkPeerFunc func(*Peer) bool) {
	for {
		peer.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		messageHeader, err := recvHeader(peer.conn)
		if nil != err {
			if io.EOF == err {
				break
			}
			_, ok := err.(*readHeaderError)
			if ok {
				time.Sleep(time.Second * 1)
				break
			}

			serr, ok := err.(net.Error)
			if ok && !serr.Timeout() {
				break
			}

			if checkPeerFunc != nil {
				if !checkPeerFunc(peer) {
					break
				}
			}

			continue
		}

		peer.Ping = time.Now()

		switch messageHeader.messageType {
		case mtProtobuf:
			msg, err := recvProtobufBody(peer.conn, *messageHeader)
			if nil != err {
				if io.EOF == err {
					break
				}
			}
			callbackProtoMsg(peer, msg)
		case mtRawbyte:
			body, err := recvRawBody(peer.conn, *messageHeader)
			if nil != err {
				if io.EOF == err {
					break
				}
			}
			callbackRawbyte(peer, messageHeader.packetType, body)
		default:
		}
	}
}
