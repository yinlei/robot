package common

import (
	"fmt"
)

type Bus chan Message

func (self *Bus) SendMessage(message *Message) {
	fmt.Println("SendMessage")
}
