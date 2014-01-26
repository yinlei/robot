package main

import (
	"fmt"
	//"time"
	//"net"
	"os"
	"os/signal"
	"syscall"
)

import (
	//"logger"
	"common"
	_ "gate"
)

func main() {
	//logger.Debugf("Debug")
	fmt.Println("System Start...")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	common.B.Go()
	<-sigs

	fmt.Println("System End...")
}
