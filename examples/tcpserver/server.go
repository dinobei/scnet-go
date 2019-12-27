package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dinobei/scnet-go"
	"github.com/dinobei/scnet-go/examples/example"
)

func main() {
	port := flag.Int("port", 8000, "service port")
	flag.Parse()

	scnet.RegistProtoMessage(0, example.Packet1{}, onPacket1)
	scnet.RegistProtoMessage(1, example.Packet2{}, onPacket2)
	scnet.RegistProtoMessage(2, example.Packet3{}, onPacket3)
	scnet.RegistProtoMessage(3, example.Packet4{}, onPacket4)
	scnet.RegistProtoMessage(4, example.ArrayMessage{}, onArrayMessage)
	scnet.RegistProtoMessage(5, example.ImageRequest{}, onImageRequest)
	scnet.RegistRawbyte(0, onRawbyte1)
	scnet.RegistRawbyte(1, onRawbyte2)

	// go scnet.Start(*port)
	server := scnet.TCPServer{}
	server.Delegate.ServerStarted = func(listener net.Listener) {
		log.Println("Server started")
	}
	server.Delegate.ServerStopped = func() {
		log.Println("Server stopped")
	}

	server.Delegate.ClientConnected = func(conn net.Conn) {
		log.Println("Client connected,", conn.RemoteAddr())
	}
	server.Delegate.ClientDisconnected = func(conn net.Conn) {
		log.Println("Client disconnected,", conn.RemoteAddr())
	}

	server.Start(*port)
	fmt.Scanln()
}

func onPacket1(conn net.Conn, data interface{}) {
	message := data.(*example.Packet1)
	log.Println("onPacket1() called, ", message.Number)

	scnet.SendProtobuf(conn, message)
}

func onPacket2(conn net.Conn, data interface{}) {
	message := data.(*example.Packet2)
	log.Println("onPacket2() called, ", message.Str)

	scnet.SendProtobuf(conn, message)
}

func onPacket3(conn net.Conn, data interface{}) {
	message := data.(*example.Packet3)
	log.Println("onPacket3() called, ", message.BoolValue)

	scnet.SendProtobuf(conn, message)
}

func onPacket4(conn net.Conn, data interface{}) {
	message := data.(*example.Packet4)
	log.Println("onPacket4() called, ", message.FloatValue, message.DoubleValue)

	scnet.SendProtobuf(conn, message)
}

func onArrayMessage(conn net.Conn, data interface{}) {
	message := data.(*example.ArrayMessage)
	log.Println("onArrayMessage() called, ", message)

	scnet.SendProtobuf(conn, message)
}

func onImageRequest(conn net.Conn, data interface{}) {
	message := data.(*example.ImageRequest)
	log.Println("onImageRequest() called, ", message)

	scnet.SendProtobuf(conn, message)
}

func onRawbyte1(conn net.Conn, buf []byte) {
	log.Printf("onRawbyte1() called, %s\n", buf)

	scnet.Send(conn, 0, buf)
}

func onRawbyte2(conn net.Conn, buf []byte) {
	log.Printf("onRawbyte2() called, %s\n", buf)

	scnet.Send(conn, 1, buf)
}
