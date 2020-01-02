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
	clients  map[string]net.Conn
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

	s.clients = make(map[string]net.Conn)

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

		s.clients[conn.RemoteAddr().String()] = conn
		s.Delegate.ClientConnected(conn)
		go s.handler(conn)
	}

	s.Delegate.ServerStopped()
}

func (s TCPServer) handler(conn net.Conn) {
	connHandler(conn)
	s.Delegate.ClientDisconnected(conn)
	delete(s.clients, conn.RemoteAddr().String())
}

func (s TCPServer) GetClient(address string) (net.Conn, bool) {
	if val, ok := s.clients[address]; ok {
		return val, true
	}
	return nil, false
}
