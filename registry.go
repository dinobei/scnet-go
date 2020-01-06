package scnet

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

// reference: https://stackoverflow.com/a/23031445
var typeRegistry = make(map[int32]reflect.Type)
var funcRegistry = make(map[string]func(*Peer, interface{}))
var packetTypeRegistry = make(map[string]int32)

var funcRawRegistry = make(map[int32]func(*Peer, []byte))

var protoMessageTypes []interface{}

// RegistProtoMessage ...
func RegistProtoMessage(packetType int32, message interface{}, f func(*Peer, interface{})) {
	typeRegistry[packetType] = reflect.TypeOf(message)
	funcRegistry[fmt.Sprintf("%T", message)] = f
	packetTypeRegistry[fmt.Sprintf("%T", message)] = packetType
	protoMessageTypes = append(protoMessageTypes, message)
}

// RegistRawbyte ...
func RegistRawbyte(packetType int32, f func(*Peer, []byte)) {
	funcRawRegistry[packetType] = f
}

func makeInstance(messageType int32) interface{} {
	if typeRegistry[messageType] == nil {
		return nil
	}

	v := reflect.New(typeRegistry[messageType])

	return v.Interface()
}

// callbackProtoMsg ...
// reference: https://stackoverflow.com/a/39144290
func callbackProtoMsg(peer *Peer, msg proto.Message) bool {
	for _, data := range protoMessageTypes {
		if proto.MessageName(msg) == reflect.TypeOf(data).String() {
			name := fmt.Sprintf("%T", data)
			f := funcRegistry[name]
			if f != nil {
				f(peer, msg)
			}

			return true
		}
	}

	return false
}

// callbackRawbyte ...
func callbackRawbyte(peer *Peer, packetType int32, buf []byte) bool {
	f := funcRawRegistry[packetType]
	if f == nil {
		return false
	}

	f(peer, buf)
	return true
}

// getPacketType : protobuf pakcet type getter
func getPacketType(name string) int32 {
	return packetTypeRegistry[name]
}
