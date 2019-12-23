package scnet

import (
	"fmt"
	"net"
	"reflect"

	"github.com/golang/protobuf/proto"
)

// reference: https://stackoverflow.com/a/23031445
var typeRegistry = make(map[int32]reflect.Type)
var funcRegistry = make(map[string]func(net.Conn, interface{}))
var packetTypeRegistry = make(map[string]int32)

var funcRawRegistry = make(map[int32]func(net.Conn, []byte))

var protoMessageTypes []interface{}

// RegistProtoMessage ...
func RegistProtoMessage(packetType int32, message interface{}, f func(net.Conn, interface{})) {
	typeRegistry[packetType] = reflect.TypeOf(message)
	funcRegistry[fmt.Sprintf("%T", message)] = f
	packetTypeRegistry[fmt.Sprintf("%T", message)] = packetType
	protoMessageTypes = append(protoMessageTypes, message)
}

// RegistRawbyte ...
func RegistRawbyte(packetType int32, f func(net.Conn, []byte)) {
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
func callbackProtoMsg(conn net.Conn, msg proto.Message) bool {
	for _, data := range protoMessageTypes {
		if proto.MessageName(msg) == reflect.TypeOf(data).String() {
			name := fmt.Sprintf("%T", data)
			f := funcRegistry[name]
			if f != nil {
				f(conn, msg)
			}

			return true
		}
	}

	return false
}

// callbackRawbyte ...
func callbackRawbyte(conn net.Conn, packetType int32, buf []byte) bool {
	f := funcRawRegistry[packetType]
	if f == nil {
		return false
	}

	f(conn, buf)
	return true
}

// getPacketType : protobuf pakcet type getter
func getPacketType(name string) int32 {
	return packetTypeRegistry[name]
}
