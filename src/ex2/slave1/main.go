package main

import (
	"log"
	"net"
	"os"
)

func main() {
	l := log.New(os.Stdout, "slave1: ", log.Ltime|log.Lmicroseconds)
	conn_socket := os.NewFile(uintptr(3), "conn")
	conn, _ := net.FileConn(conn_socket)

	l.Printf("pid %d started\n", os.Getpid())

	defer conn_socket.Close()
	defer conn.Close()
	defer l.Println("Connection closed")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		conn.Write(buf[:n])
	}
}
