package main

import (
	"log"
	"net"
	"os"
)

func main() {
	l := log.New(os.Stdout, "ex1: ", log.Ltime|log.Lmicroseconds)

	addr := net.TCPAddr{Port: 6000}
	ln, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		l.Printf("Fail to listen: %s", err)
		os.Exit(1)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			l.Printf("Fail to accept: %s", err)
			os.Exit(1)
		}

		go func(conn net.Conn) {
			l.Printf("Conn from %s", conn.RemoteAddr())
			buf := make([]byte, 1024)

			defer conn.Close()
			defer l.Println("Connection closed")
			for {
				n, err := conn.Read(buf)
				if err != nil {
					return
				}

				conn.Write(buf[:n])
			}
		}(conn)
	}
}
