package main

import (
	"flag"
	"fmt"
	"log"

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
	server.Delegate.ServerStarted = func(addr string) {
		log.Println("Server started")
	}

	server.Delegate.ServerStopped = func() {
		log.Println("Server stopped")
	}

	server.Delegate.ClientConnected = func(peer scnet.Peer) {
		log.Println("Client connected,", peer.GetRemoteAddr())
	}

	server.Delegate.ClientDisconnected = func(peer scnet.Peer) {
		log.Println("Client disconnected,", peer.GetRemoteAddr())
	}

	server.Start(*port)
	fmt.Scanln()
}

func onPacket1(peer *scnet.Peer, data interface{}) {
	message := data.(*example.Packet1)
	log.Println("onPacket1() called, ", message.Number)

	scnet.Send(*peer, message)
}

func onPacket2(peer *scnet.Peer, data interface{}) {
	message := data.(*example.Packet2)
	log.Println("onPacket2() called, ", message.Str)

	scnet.Send(*peer, message)
}

func onPacket3(peer *scnet.Peer, data interface{}) {
	message := data.(*example.Packet3)
	log.Println("onPacket3() called, ", message.BoolValue)

	scnet.Send(*peer, message)
}

func onPacket4(peer *scnet.Peer, data interface{}) {
	message := data.(*example.Packet4)
	log.Println("onPacket4() called, ", message.FloatValue, message.DoubleValue)

	scnet.Send(*peer, message)
}

func onArrayMessage(peer *scnet.Peer, data interface{}) {
	message := data.(*example.ArrayMessage)
	log.Println("onArrayMessage() called, ", message)

	scnet.Send(*peer, message)
}

func onImageRequest(peer *scnet.Peer, data interface{}) {
	message := data.(*example.ImageRequest)
	log.Println("onImageRequest() called, ", message)

	scnet.Send(*peer, message)
}

func onRawbyte1(peer *scnet.Peer, buf []byte) {
	log.Printf("onRawbyte1() called, %s\n", buf)

	data := scnet.RawbyteData{}
	data.PacketType = 0
	data.Buffer = buf
	scnet.Send(*peer, data)
}

func onRawbyte2(peer *scnet.Peer, buf []byte) {
	log.Printf("onRawbyte2() called, %s\n", buf)

	data := scnet.RawbyteData{}
	data.PacketType = 1
	data.Buffer = buf
	scnet.Send(*peer, data)
}
