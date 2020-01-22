package main

import (
	"log"
	"net"
)

func server() {
	conn, err := net.ListenPacket("udp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1500)
	for {
		len, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Println(err)
		}
		if buf[1] == MessageTypeEchoRequest {
			log.Printf("received echo request from %s", addr)
			buf[1] = MessageTypeEchoResponse
			conn.WriteTo(buf[:len], addr)
		}
	}
}
