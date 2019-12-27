package scnet

import (
	"log"
	"net"
	"strconv"
	"time"
)

// TCPServer ...
type TCPServer struct {
	address  string
	Delegate TCPServerDelegate
}

// TCPServerDelegate ...
type TCPServerDelegate struct {
	ServerStarted      func(net.Listener)
	ServerStopped      func()
	ClientConnected    func(net.Conn)
	ClientDisconnected func(net.Conn)
}

// Start ...
func (s TCPServer) Start(port int) {
	s.address = ":" + strconv.Itoa(port)

	l, err := net.Listen("tcp", s.address)
	s.Delegate.ServerStarted(l)
	if nil != err {
		log.Fatalf("fail to bind address to %d; err: %v", port, err)
	}
	defer l.Close()

	tcplistener := l.(*net.TCPListener)
	defer tcplistener.Close()

	for {
		tcplistener.SetDeadline(time.Now().Add(time.Second * 10))

		conn, err := l.Accept()
		if nil != err {
			if err, ok := err.(*net.OpError); ok && err.Timeout() {
				continue
			}

			log.Printf("fail to accept; err: %v", err)
			break
		}
		defer conn.Close()

		s.Delegate.ClientConnected(conn)
		go s.handler(conn)
	}

	s.Delegate.ServerStopped()
}

func (s TCPServer) handler(conn net.Conn) {
	connHandler(conn)
	s.Delegate.ClientDisconnected(conn)
}
