package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dinobei/scnet-go"
	"github.com/dinobei/scnet-go/examples/example"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "server ip")
	port := flag.Int("port", 8000, "server port")
	flag.Parse()

	scnet.RegistProtoMessage(0, example.Packet1{}, onPacket1)
	scnet.RegistProtoMessage(1, example.Packet2{}, onPacket2)
	scnet.RegistProtoMessage(2, example.Packet3{}, onPacket3)
	scnet.RegistProtoMessage(3, example.Packet4{}, onPacket4)
	scnet.RegistProtoMessage(4, example.ArrayMessage{}, onArrayMessage)
	scnet.RegistProtoMessage(5, example.ImageRequest{}, onImageRequest)
	scnet.RegistRawbyte(0, onRawbyte1)
	scnet.RegistRawbyte(1, onRawbyte2)
	scnet.RegistRawbyte(2, onPing)

	client := scnet.TCPClient{}
	client.Delegate.Attached = func() {
		log.Println("Attached")
	}
	client.Delegate.Detached = func() {
		log.Println("Detached")
	}
	client.Delegate.TimedOut = func(peer *scnet.Peer) bool {
		ping := scnet.RawbyteData{}
		ping.PacketType = 2
		scnet.Send(*peer, ping)
		return true
	}
	client.Delegate.Connected = func(peer scnet.Peer) {
		log.Println("Connected, ", peer.GetRemoteAddr())

		go func() {
			for {
				pkt1 := &example.Packet1{}
				pkt1.Number = int32(rand.Intn(1000))
				err := scnet.Send(peer, pkt1)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				pkt2 := &example.Packet2{}
				pkt2.Str = "Hello World"
				err = scnet.Send(peer, pkt2)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				pkt3 := &example.Packet3{}
				pkt3.BoolValue = true
				err = scnet.Send(peer, pkt3)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				pkt4 := &example.Packet4{}
				pkt4.DoubleValue = 1.1
				pkt4.FloatValue = 2.2
				err = scnet.Send(peer, pkt4)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				arrayMessage := &example.ArrayMessage{}
				arrayMessage.StrArr = append(arrayMessage.StrArr, "sample1")
				arrayMessage.StrArr = append(arrayMessage.StrArr, "sample2")
				err = scnet.Send(peer, arrayMessage)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				imgReq := &example.ImageRequest{}
				imgReq.Name = "test.jpg"
				err = scnet.Send(peer, imgReq)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				data0 := scnet.RawbyteData{}
				data0.PacketType = 0
				data0.Buffer = []byte("raw pakcet (type 0)")
				err = scnet.Send(peer, data0)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				data1 := scnet.RawbyteData{}
				data1.PacketType = 1
				data1.Buffer = []byte("raw pakcet (type 1)")
				err = scnet.Send(peer, data1)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)
			}
		}()

	}
	client.Delegate.Connecting = func(address string) {
		log.Println("Connecting...", address)
	}
	client.Delegate.Disconnected = func(peer scnet.Peer) {
		log.Println("Disconnected, ", peer.GetRemoteAddr())
	}

	go client.Attach(*ip, *port, time.Second*5)

	fmt.Scanln()
	client.Detach()
	fmt.Scanln()
}

func onPacket1(peer *scnet.Peer, data interface{}) {
	message := data.(*example.Packet1)
	log.Println("onPacket1() called, ", message.Number)
}

func onPacket2(peer *scnet.Peer, data interface{}) {
	message := data.(*example.Packet2)
	log.Println("onPacket2() called, ", message.Str)
}

func onPacket3(peer *scnet.Peer, data interface{}) {
	message := data.(*example.Packet3)
	log.Println("onPacket3() called, ", message.BoolValue)
}

func onPacket4(peer *scnet.Peer, data interface{}) {
	message := data.(*example.Packet4)
	log.Println("onPacket4() called, ", message.FloatValue, message.DoubleValue)
}

func onArrayMessage(peer *scnet.Peer, data interface{}) {
	message := data.(*example.ArrayMessage)
	log.Println("onArrayMessage() called, ", message)
}

func onImageRequest(peer *scnet.Peer, data interface{}) {
	message := data.(*example.ImageRequest)
	log.Println("onImageRequest() called, ", message)
}

func onRawbyte1(peer *scnet.Peer, buf []byte) {
	log.Printf("onRawbyte1() called, %s\n", buf)
}

func onRawbyte2(peer *scnet.Peer, buf []byte) {
	log.Printf("onRawbyte2() called, %s\n", buf)
}

func onPing(peer *scnet.Peer, buf []byte) {
	log.Printf("onPing() called")
}
