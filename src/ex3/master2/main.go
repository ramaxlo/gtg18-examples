package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	l := log.New(os.Stdout, "ex3: ", log.Ltime|log.Lmicroseconds)

	addr := net.TCPAddr{Port: 6000}
	ln, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		l.Printf("Fail to listen: %s", err)
		os.Exit(1)
	}
	defer ln.Close()

	notify_ch := make(chan bool, 0)
	for {
		file, _ := ln.File()
		fd, _ := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
		f1 := os.NewFile(uintptr(fd[0]), "socket1")
		f2 := os.NewFile(uintptr(fd[1]), "socket2")

		cmd := exec.Command("slave2")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.ExtraFiles = append(cmd.ExtraFiles, file, f2)
		cmd.Start()

		// Must close only after child process is started
		file.Close()
		f2.Close()

		go func() {
			defer f1.Close()

			buf := make([]byte, 16)
			_, err := f1.Read(buf)
			if err != nil {
				l.Printf("Fail to read socket: %s\n", err)
			}

			notify_ch <- true

			// Must wait for child's termination
			if err := cmd.Wait(); err != nil {
				l.Printf("Error state: %s", err)
			}
		}()

		<-notify_ch
	}
}
