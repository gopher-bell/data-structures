// TCP Client
package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5032")
	if nil != err {
		log.Fatalf("failed to connect to server")
	}

	// some event happens
	s := make([]byte, 50000)
	for i := range s {
		s[i] = 'A'
	}

	for {
		// heartbeat
		n, err := conn.Write(s)
		log.Println(n)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Duration(3) * time.Second)
	}
}
