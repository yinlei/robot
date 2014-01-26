package agent

import (
	"net"
)

type Agent struct {
	id int
	conn net.Conn
	incoming chan []byte
	outgoing chan []byte
	close chan bool
}

var id int = 0

func readHander(a *Agent) {
	buffer := make([]byte, 1024)

	for a.Read(buffer) {
		a.incoming <- buffer
	}
}

func writeHandler(a *Agent) {
	for {
		select {
		case buffer <- a.outgoing:
			a.Write(buffer)
		case <- a.close:
			a.conn.Close()
			break
		}
	}
}

func NewAgent(conn net.Conn) *Agent {
	id += 1
	agent := &Agent{id, conn, make(chan []byte), make(chan []byte), make(chan bool)}
	go readHandler(agent)
	go writeHandler(agent)
	return  agent
}

func (self *Agent) Read(buffer []byte) bool {
	bytesRead, err := net.Conn.Read(buffer)
	if err != nil {
		self.Close()
		return false
	}

	return true
}

func (self *Agent) Write(buffer []byte) bool {
	bytesRead, err := net.Conn.Write(buffer)

	if err != nil {
		return false
	}

	return true
}

func (self *Agent) Close() {
	self.close <- true
	self.conn.Close()
}

func (self *Agent) Equal(dst *Agent) bool {
	return false
}

func (self *Agent) SendBuffer(buffer []byte) {
	self.outgoing <- buffer
}


