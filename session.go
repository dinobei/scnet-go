package scnet

import (
	"net"
	"time"
)

type readHeaderError struct{}

func (m *readHeaderError) Error() string {
	return "failed reading header"
}

type unknownProtobufPacketError struct{}

func (m *unknownProtobufPacketError) Error() string {
	return "unknown protobuf packet"
}

type notSupportedDataError struct{}

func (m *notSupportedDataError) Error() string {
	return "not supported data"
}

// Peer ...
type Peer struct {
	conn net.Conn
	Ping time.Time
}

// GetRemoteAddr ...
func (peer Peer) GetRemoteAddr() string {
	return peer.conn.RemoteAddr().String()
}

// GetLocalAddr ...
func (peer Peer) GetLocalAddr() string {
	return peer.conn.LocalAddr().String()
}

// GetFormattedPing ...
func (peer Peer) GetFormattedPing() string {
	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := peer.Ping.In(loc)
	return kst.Format("2006-01-02 15:04:05")
}
