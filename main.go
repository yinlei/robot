//
package main

import  "fmt"
import "log"
import "bytes"
import "net"
import "encoding/binary"

func main() {
	l, err := net.Listen("tcp", ":8700")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(conn.RemoteAddr())
		go func(conn net.Conn) {
			var hello = []byte{'h', 'e'}
			sendCount, err := conn.Write(hello)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(sendCount)
			for {
				var buf = make([]byte, 1000)
				nread, err := conn.Read(buf)
				if err != nil {
					log.Fatal(err)
				}
				header := bytes.NewBuffer(buf[2:4])
				var length uint16
				binary.Read(header, binary.LittleEndian, &length)
				fmt.Println(length)

				fmt.Println(nread)
			}
		}(conn)
	}
}
