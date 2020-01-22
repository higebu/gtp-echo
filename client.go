package main

import (
	"encoding/binary"
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
	var seq uint16
	b := []byte{0x32, 0x1, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	for {
		binary.BigEndian.PutUint16(b[8:10], seq)
		conn.Write(b)

		buf := make([]byte, 1500)
		_, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
		}
		if buf[1] == MessageTypeEchoResponse {
			log.Printf("received echo response")
		}
		seq += 1
		time.Sleep(1 * time.Second)
	}
}
