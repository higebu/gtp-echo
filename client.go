package main

import (
	"log"
	"net"
	"time"
)

func client() {
	conn, err := net.Dial("udp", *dstAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		conn.Write([]byte{0x32, 0x1, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x40, 0x15, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})

		buf := make([]byte, 1500)
		_, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
		}
		if buf[1] == MessageTypeEchoResponse {
			log.Printf("received echo response")
		}
		time.Sleep(1 * time.Second)
	}
}
