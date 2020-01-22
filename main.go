package main

import (
	"flag"
	"time"
)

var (
	listenAddr = flag.String("listen-addr", "127.0.0.1:2152", "listen address for server")
	dstAddr    = flag.String("dst-addr", "127.0.0.1:2152", "destination address for client")
	clientFlag = flag.Bool("client", false, "run as a client")
	interval   = flag.Duration("interval", 1*time.Second, "send interval for client")
)

func main() {
	flag.Parse()
	if *clientFlag {
		client()
	} else {
		server()
	}
}
