package scnet

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

// ClientCallback ...
type ClientCallback func(*TCPClient, Header, proto.Message)

// PeerCallback ...
type PeerCallback func(*Peer, Header, proto.Message)

// reference: https://stackoverflow.com/a/23031445
var typeRegistry = make(map[uint32]reflect.Type)
var clientCallbackRegistry = make(map[string]ClientCallback)
var peerCallbackRegistry = make(map[string]PeerCallback)
var packetTypeRegistry = make(map[string]uint32)

var protoMessageTypes []interface{}

func registMessage(packetType uint32, message interface{}, f ClientCallback) {
	typeRegistry[packetType] = reflect.TypeOf(message)
	clientCallbackRegistry[fmt.Sprintf("%T", message)] = f
	packetTypeRegistry[fmt.Sprintf("%T", message)] = packetType
	protoMessageTypes = append(protoMessageTypes, message)
}

func registMessageWithPeer(packetType uint32, message interface{}, f PeerCallback) {
	typeRegistry[packetType] = reflect.TypeOf(message)
	peerCallbackRegistry[fmt.Sprintf("%T", message)] = f
	packetTypeRegistry[fmt.Sprintf("%T", message)] = packetType
	protoMessageTypes = append(protoMessageTypes, message)
}

func makeInstance(packetType uint32) proto.Message {
	if typeRegistry[packetType] == nil {
		return nil
	}

	v := reflect.New(typeRegistry[packetType])

	return v.Interface().(proto.Message)
}

// reference: https://stackoverflow.com/a/39144290
func clientCallbackMessage(client *TCPClient, header Header, msg proto.Message) bool {
	for _, data := range protoMessageTypes {
		if proto.MessageName(msg) == reflect.TypeOf(data).String() {
			name := fmt.Sprintf("%T", data)
			f := clientCallbackRegistry[name]
			if f != nil {
				f(client, header, msg)
			}

			return true
		}
	}

	return false
}

func peerCallbackMessage(peer *Peer, header Header, msg proto.Message) bool {
	for _, data := range protoMessageTypes {
		if proto.MessageName(msg) == reflect.TypeOf(data).String() {
			name := fmt.Sprintf("%T", data)
			f := peerCallbackRegistry[name]
			if f != nil {
				f(peer, header, msg)
			}

			return true
		}
	}

	return false
}

// getPacketType : protobuf pakcet type getter
func getPacketType(name string) uint32 {
	return packetTypeRegistry[name]
}
