package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/dinobei/scnet-go"
	"github.com/dinobei/scnet-go/examples/example"
	"github.com/golang/protobuf/proto"
)

var server scnet.TCPServer

func main() {
	port := flag.Int("port", 9190, "service port")
	flag.Parse()

	server = scnet.TCPServer{}
	server.RegistMessage(uint32(example.PacketType_dummyPacket1), example.DummyPacket1{}, onDummyPacket1)
	server.RegistMessage(uint32(example.PacketType_dummyPacket2), example.DummyPacket2{}, onDummyPacket2)
	server.RegistMessage(uint32(example.PacketType_ping), example.Ping{}, onPing)

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

	server.Delegate.ClientTimeout = func(peer scnet.Peer) bool {
		current := time.Now()
		if current.Sub(peer.Ping).Seconds() > 15 {
			log.Println("ClientTimeout, exceed 15sec")
			return false
		}

		return true
	}

	go server.Start(*port)
	fmt.Scanln()
}

func onDummyPacket1(peer *scnet.Peer, header scnet.Header, message proto.Message) {
	pkt1 := message.(*example.DummyPacket1)
	log.Println("onDummyPacket1() called, ", pkt1.Number)

	server.Send(peer, &header, pkt1)
}

func onDummyPacket2(peer *scnet.Peer, header scnet.Header, message proto.Message) {
	pkt2 := message.(*example.DummyPacket2)
	log.Println("onDummyPacket2() called, StrArrLength: ", len(pkt2.StrArr))

	server.Send(peer, &header, message)
}

func onPing(peer *scnet.Peer, header scnet.Header, message proto.Message) {
	ping := message.(*example.Ping)
	server.Send(peer, &header, ping)
}
