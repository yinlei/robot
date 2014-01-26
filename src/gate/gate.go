package gate

import (
	"fmt"
	"net"
	"time"
)

import (
	"common"
	//"logger"
)

type Gate struct {
	listener net.Listener
	conns    chan net.Conn
	bus      chan common.Message
}

func init() {
	fmt.Println("Gate init")
	g := newGate()
	g.Start("")
	common.B.Regist(g)
}

func newGate() common.Service {
	return &Gate{conns: make(chan net.Conn, 10000), bus: make(chan common.Message, 10)}
}

func (self *Gate) Start(name string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8081")
	fmt.Println(tcpAddr)
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		//logger.Errorf(err)
		fmt.Println(err)
	}
	fmt.Println("监听启动")

	//self.listener = l
	for i := 1; i < 10; i++ {
		go self.doAccept(l)
		fmt.Println("启动一个Accept")
	}

	//go self.Loop()
}

func (self *Gate) doAccept(l *net.TCPListener) {
	for {
		_, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			//self.conns <- conn
			fmt.Println("Tcp Accept")
		}
	}
}

func (self *Gate) Loop() {
	for {
		select {
		case conn := <-self.conns:
			fmt.Println(conn.LocalAddr().String())
			fmt.Println("创建一个连接")
		case <-time.After(time.Second * 10):
			fmt.Println("超时")
		}
	}
}

func (self *Gate) Stop() {
	self.listener.Close()
}

func (self *Gate) Name() string {
	return "Gate"
}

func (self *Gate) Call(name string, params ...interface{}) bool {
	return false
}
