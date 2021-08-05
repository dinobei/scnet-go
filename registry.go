package scnet

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
)

// ClientCallback ...
type ClientCallback func(*TCPClient, Header, proto.Message)

// PeerCallback ...
type PeerCallback func(*Peer, Header, proto.Message)

// reference: https://stackoverflow.com/a/23031445
var typeRegistry = make(map[string]reflect.Type)
var clientCallbackRegistry = make(map[string]ClientCallback) // TODO: used in client
var peerCallbackRegistry = make(map[string]PeerCallback) // TODO: used in server


func registMessage(message interface{}, f ClientCallback) {
	packetType := strings.ToLower(fmt.Sprintf("%T", message))
	typeRegistry[packetType] = reflect.TypeOf(message)
	clientCallbackRegistry[packetType] = f
}

func registMessageWithPeer(message interface{}, f PeerCallback) {
	packetType := strings.ToLower(fmt.Sprintf("%T", message))
	typeRegistry[packetType] = reflect.TypeOf(message)
	peerCallbackRegistry[packetType] = f
}

func makeInstance(packetType string) proto.Message {
	if typeRegistry[packetType] == nil {
		return nil
	}

	v := reflect.New(typeRegistry[packetType])

	return v.Interface().(proto.Message)
}

// reference: https://stackoverflow.com/a/39144290
func clientCallbackMessage(client *TCPClient, header Header, msg proto.Message) bool {
	f, exists := clientCallbackRegistry[getPacketType(msg)]
	if !exists {
		return false;
	}
	if f != nil {
		f(client, header, msg)
	}
	return true
}

func peerCallbackMessage(peer *Peer, header Header, msg proto.Message) bool {
	f, exists := peerCallbackRegistry[getPacketType(msg)]
	if !exists {
		return false;
	}
	f(peer, header, msg)
	return true
}

func getPacketType(msg proto.Message) string {
	var packetType = fmt.Sprintf("%T", msg)
	packetType = strings.ToLower(packetType)
	packetType = strings.TrimLeft(packetType, "*")
	return packetType
}
