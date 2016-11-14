package main

import (
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	l := log.New(os.Stdout, "ex2: ", log.Ltime|log.Lmicroseconds)

	addr := net.TCPAddr{Port: 6000}
	ln, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		l.Printf("Fail to listen: %s", err)
		os.Exit(1)
	}
	defer ln.Close()

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			l.Printf("Fail to accept: %s", err)
			os.Exit(1)
		}

		l.Printf("Conn from %s", conn.RemoteAddr())

		file, _ := conn.File()
		cmd := exec.Command("slave1")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.ExtraFiles = append(cmd.ExtraFiles, file)
		cmd.Start()

		// Must close only after child process is started
		file.Close()
		conn.Close()

		go func() {
			// Must wait for child's termination
			if err := cmd.Wait(); err != nil {
				l.Printf("Error state: %s", err)
			}
		}()
	}
}
