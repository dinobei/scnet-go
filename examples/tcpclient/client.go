package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dinobei/scnet-go"
	"github.com/dinobei/scnet-go/examples/example"
	"github.com/golang/protobuf/proto"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "server ip")
	port := flag.Int("port", 9190, "server port")
	flag.Parse()

	client := scnet.NewTCPClient()
	client.RegistMessage(uint32(example.PacketType_dummyPacket1), example.DummyPacket1{}, onDummyPacket1)
	client.RegistMessage(uint32(example.PacketType_dummyPacket2), example.DummyPacket2{}, onDummyPacket2)
	client.RegistMessage(uint32(example.PacketType_ping), example.Ping{}, onPing)
	client.Delegate.Attached = func() {
		log.Println("Attached")
	}
	client.Delegate.Detached = func() {
		log.Println("Detached")
	}
	client.Delegate.TimedOut = func(client *scnet.TCPClient) bool {
		ping := &example.Ping{}
		client.Send(nil, ping, nil)
		return true
	}
	client.Delegate.Connected = func(client *scnet.TCPClient) {
		log.Println("Connected, ", client.GetAddress())

		go func() {
			for {
				pkt1 := &example.DummyPacket1{}
				pkt1.Number = int32(rand.Intn(1000))
				pkt1.Title = "Hello World"
				if !client.Send(nil, pkt1,
					func(client *scnet.TCPClient, header scnet.Header, message proto.Message) {
						dummyPacket1 := message.(*example.DummyPacket1)
						log.Print("[DedicatedCallback] onPacket1() called, Number:", dummyPacket1.Number, ", Str:", dummyPacket1.Title)
					}) {
					break
				}
				time.Sleep(time.Second * 1)

				pkt2 := &example.DummyPacket2{}
				pkt2.StrArr = append(pkt2.StrArr, "This")
				pkt2.StrArr = append(pkt2.StrArr, "is")
				pkt2.StrArr = append(pkt2.StrArr, "sample")
				pkt2.StrArr = append(pkt2.StrArr, "message")
				if !client.Send(nil, pkt2, nil) {
					break
				}
				time.Sleep(time.Second * 1)
			}
		}()

	}
	client.Delegate.Connecting = func(client *scnet.TCPClient) {
		log.Println("Connecting...", client.GetAddress())
	}
	client.Delegate.Disconnected = func(client *scnet.TCPClient) {
		log.Println("Disconnected", client.GetAddress())
	}

	go client.Attach(*ip, *port, time.Second*5)

	fmt.Scanln()
	client.Detach()
	fmt.Scanln()
}

func onDummyPacket1(client *scnet.TCPClient, header scnet.Header, message proto.Message) {
	dummyPacket1 := message.(*example.DummyPacket1)
	log.Print("[SinkCallback] onPacket1() called, Number:", dummyPacket1.Number, ", Str:", dummyPacket1.Title)
}

func onDummyPacket2(client *scnet.TCPClient, header scnet.Header, message proto.Message) {
	dummyPacket2 := message.(*example.DummyPacket2)

	strArr := make([]string, 0, len(dummyPacket2.StrArr))
	for _, str := range dummyPacket2.StrArr {
		strArr = append(strArr, str)
	}
	log.Print("[SinkCallback] onPacket2() called, StrArrLength:", len(dummyPacket2.StrArr), ", StrArr: ", strArr)
}

func onPing(client *scnet.TCPClient, header scnet.Header, message proto.Message) {
	log.Println("[SinkCallback] onPing() called")
}
