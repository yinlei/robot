package common

import (
	"fmt"
)

type Board struct {
	bus   Bus
	slots map[string]interface{}
}

var B *Board = &Board{slots: make(map[string]interface{})}

//	注册服务
func (self *Board) Regist(service interface{}) {
	fmt.Println("Regist")
	self.slots[service.(Service).Name()] = service
	fmt.Println(service.(Service).Name())
}

func (self *Board) Dispatch() {
	m := <-self.bus
	fmt.Println("Dispatch Message")
	fmt.Println(m.id)
}

func (self *Board) Go() {
	go self.Dispatch()
}
