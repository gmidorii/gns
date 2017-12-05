package main

import (
	"fmt"
	"log"
	"net"
)

func run() error {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 53,
	}
	l, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer l.Close()

	buf := make([]byte, 1024)
	for {
		_, addr, err := l.ReadFromUDP(buf)
		if err != nil {
			return err
		}

		go dns(l, addr, buf)
	}
}

func dns(l *net.UDPConn, addr *net.UDPAddr, buf []byte) {
	fmt.Println(string(buf))
	l.WriteToUDP([]byte("PONG"), addr)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
