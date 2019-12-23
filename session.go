package scnet

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
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
	responseSize := int(header.dataSize)

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

	msg := makeInstance(header.packetType)
	if msg == nil {
		println("unknown protobuf packet")
		return nil, &unknownProtobufPacketError{}
	}

	field := msg.(proto.Message)

	err := proto.Unmarshal(buf, field)
	if err != nil {
		return nil, err
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

// SendProtobuf ...
func SendProtobuf(conn net.Conn, message proto.Message) error {
	packetType := getPacketType(proto.MessageName(message))

	buffer := makeHeader(mtProtobuf, packetType, proto.Size(message))

	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	buffer = append(buffer, data[:]...)

	_, err = conn.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}

// Send ...
func Send(conn net.Conn, packetType int32, data []byte) error {
	buffer := makeHeader(mtRawbyte, packetType, len(data))
	buffer = append(buffer, data[:]...)

	_, err := conn.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}

// Start ...
func Start(port int) {
	address := ":" + strconv.Itoa(port)
	fmt.Println("service started:", address)

	l, err := net.Listen("tcp", address)
	if nil != err {
		log.Fatalf("fail to bind address to %d; err: %v", port, err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Printf("fail to accept; err: %v", err)
			continue
		}

		log.Printf("connected client, %s", conn.RemoteAddr())
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		messageHeader, err := recvHeader(conn)
		if nil != err {
			if io.EOF == err {
				break
			}

			fmt.Println(err)
			continue
		}

		switch messageHeader.messageType {
		case mtProtobuf:
			msg, err := recvProtobufBody(conn, *messageHeader)
			if nil != err {
				if io.EOF == err {
					break
				}
			}

			callbackProtoMsg(conn, msg)

		case mtRawbyte:
			body, err := recvRawBody(conn, *messageHeader)
			if nil != err {
				if io.EOF == err {
					break
				}
			}

			callbackRawbyte(conn, messageHeader.packetType, body)

		default:

		}

	}

	log.Printf("disconnected client, %s", conn.RemoteAddr())
}

// Conn ...
func Conn(ip string, port int) net.Conn {
	addr := ip + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go connHandler(conn)
	return conn
}