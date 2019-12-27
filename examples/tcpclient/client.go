package main

import (
	"flag"
	"fmt"
	"github.com/dinobei/scnet-go"
	"github.com/dinobei/scnet-go/examples/example"
	"log"
	"math/rand"
	"net"
	"time"
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

	client := scnet.TCPClient{}
	client.Delegate.Attached = func() {
		log.Println("Attached")
	}
	client.Delegate.Detached = func() {
		log.Println("Detached")
	}
	client.Delegate.Connected = func(conn net.Conn) {
		log.Println("Connected, ", conn.RemoteAddr())

		go func() {
			for {
				pkt1 := &example.Packet1{}
				pkt1.Number = int32(rand.Intn(1000))
				err := scnet.SendProtobuf(conn, pkt1)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				pkt2 := &example.Packet2{}
				pkt2.Str = "Hello World"
				err = scnet.SendProtobuf(conn, pkt2)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				pkt3 := &example.Packet3{}
				pkt3.BoolValue = true
				err = scnet.SendProtobuf(conn, pkt3)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				pkt4 := &example.Packet4{}
				pkt4.DoubleValue = 1.1
				pkt4.FloatValue = 2.2
				err = scnet.SendProtobuf(conn, pkt4)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				arrayMessage := &example.ArrayMessage{}
				arrayMessage.StrArr = append(arrayMessage.StrArr, "sample1")
				arrayMessage.StrArr = append(arrayMessage.StrArr, "sample2")
				err = scnet.SendProtobuf(conn, arrayMessage)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				imgReq := &example.ImageRequest{}
				imgReq.Name = "test.jpg"
				err = scnet.SendProtobuf(conn, imgReq)
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				err = scnet.Send(conn, 0, []byte("raw packet (type 0)"))
				if err != nil {
					break
				}
				time.Sleep(time.Second * 1)

				err = scnet.Send(conn, 1, []byte("raw packet (type 1)"))
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
	client.Delegate.Disconnected = func(conn net.Conn) {
		log.Println("Disconnected, ", conn.RemoteAddr())
	}

	go client.Attach(*ip, *port, time.Second*5)

	fmt.Scanln()
	client.Detach()
	fmt.Scanln()
}

func onPacket1(conn net.Conn, data interface{}) {
	message := data.(*example.Packet1)
	log.Println("onPacket1() called, ", message.Number)
}

func onPacket2(conn net.Conn, data interface{}) {
	message := data.(*example.Packet2)
	log.Println("onPacket2() called, ", message.Str)
}

func onPacket3(conn net.Conn, data interface{}) {
	message := data.(*example.Packet3)
	log.Println("onPacket3() called, ", message.BoolValue)
}

func onPacket4(conn net.Conn, data interface{}) {
	message := data.(*example.Packet4)
	log.Println("onPacket4() called, ", message.FloatValue, message.DoubleValue)
}

func onArrayMessage(conn net.Conn, data interface{}) {
	message := data.(*example.ArrayMessage)
	log.Println("onArrayMessage() called, ", message)
}

func onImageRequest(conn net.Conn, data interface{}) {
	message := data.(*example.ImageRequest)
	log.Println("onImageRequest() called, ", message)
}

func onRawbyte1(conn net.Conn, buf []byte) {
	log.Printf("onRawbyte1() called, %s\n", buf)
}

func onRawbyte2(conn net.Conn, buf []byte) {
	log.Printf("onRawbyte2() called, %s\n", buf)
}
