package scnet

import (
	"net"
	"strconv"
	"sync"
	"time"
)

// TCPClient ...
type TCPClient struct {
	address  string
	Delegate TCPClientDelegate

	conn net.Conn

	isStop bool
	mutex  sync.RWMutex
}

// TCPClientDelegate ...
type TCPClientDelegate struct {
	Attached     func()
	Detached     func()
	Connected    func(net.Conn)
	Connecting   func(string)
	Disconnected func(net.Conn)
}

// Attach ...
func (c *TCPClient) Attach(ip string, port int, timeout time.Duration) {
	c.address = ip + ":" + strconv.Itoa(port)

	if c.Delegate.Attached != nil {
		c.Delegate.Attached()
	}

	for {
		c.mutex.RLock()
		if c.isStop {
			break
		}
		c.mutex.RUnlock()

		c.conn, _ = net.DialTimeout("tcp", c.address, timeout)

		if c.conn == nil {
			// timeout
			if c.Delegate.Connecting != nil {
				c.Delegate.Connecting(c.address)
			}

			time.Sleep(timeout)
			continue
		}

		if c.Delegate.Connected != nil {
			c.Delegate.Connected(c.conn)
		}

		connHandler(c.conn)
		if c.Delegate.Disconnected != nil {
			c.Delegate.Disconnected(c.conn)
		}

		c.conn.Close()
		c.conn = nil
	}

	if c.Delegate.Detached != nil {
		c.Delegate.Detached()
	}
}

// Detach ...
func (c *TCPClient) Detach() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isStop {
		return
	}
	c.isStop = true

	if c.conn != nil {
		c.conn.Close()
	}
}
