package main

import (
	"log"
	"net"
	"os"
)

func main() {
	l := log.New(os.Stdout, "slave2: ", log.Ltime|log.Lmicroseconds)
	ln_socket := os.NewFile(uintptr(3), "listener")
	ln, _ := net.FileListener(ln_socket)
	notify := os.NewFile(uintptr(4), "notify")

	l.Printf("pid %d started\n", os.Getpid())

	defer notify.Close()
	defer ln_socket.Close()
	defer ln.Close()
	defer l.Println("Connection closed")

	conn, err := ln.Accept()
	if err != nil {
		l.Printf("Fail to accept: %s", err)
		return
	}
	defer conn.Close()

	l.Printf("Conn from %s\n", conn.RemoteAddr())
	notify.Write([]byte("gotit"))

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		conn.Write(buf[:n])
	}
}
