package scnet

import (
	"log"
	"net"
	"strconv"
	"time"
)

var clients map[string]*Peer

// TCPServer ...
type TCPServer struct {
	address  string
	Delegate TCPServerDelegate
}

// TCPServerDelegate ...
type TCPServerDelegate struct {
	ServerStarted      func(string)
	ServerStopped      func()
	ClientConnected    func(Peer)
	ClientDisconnected func(Peer)
}

// Start ...
func (s TCPServer) Start(port int) {
	s.address = ":" + strconv.Itoa(port)

	l, err := net.Listen("tcp", s.address)
	if s.Delegate.ServerStarted != nil {
		s.Delegate.ServerStarted(l.Addr().String())
	}

	if nil != err {
		log.Fatalf("fail to bind address to %d; err: %v", port, err)
	}
	defer l.Close()

	clients = make(map[string]*Peer)

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

		var peer *Peer
		if p, ok := clients[conn.RemoteAddr().String()]; ok {
			peer = p
		} else {
			peer = &Peer{}
			peer.conn = conn
		}
		peer.Ping = time.Now()

		clients[conn.RemoteAddr().String()] = peer

		if s.Delegate.ClientConnected != nil {
			s.Delegate.ClientConnected(*peer)
		}

		go s.handler(peer)
	}

	if s.Delegate.ServerStopped != nil {
		s.Delegate.ServerStopped()
	}
}

func (s TCPServer) handler(peer *Peer) {
	connHandler(peer, s.CheckPeer)

	if s.Delegate.ClientDisconnected != nil {
		s.Delegate.ClientDisconnected(*peer)
	}

	delete(clients, peer.conn.RemoteAddr().String())

	time.Sleep(time.Second * 3)
	defer peer.conn.Close()
}

// GetClient ...
func (s TCPServer) GetClient(address string) (*Peer, bool) {
	if val, ok := clients[address]; ok {
		return val, true
	}
	return nil, false
}

// CheckPeer ...
func (s TCPServer) CheckPeer(peer *Peer) bool {

	current := time.Now()
	if current.Sub(peer.Ping).Seconds() > 15 {
		return false
	}

	return true
}
