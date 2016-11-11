package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func usage() {
	fmt.Println("Usage: ./client <addr> <num of connections>")
}

func start(l *log.Logger, idx int, wg *sync.WaitGroup, addr string) {
	defer wg.Done()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		l.Printf("%d: Fail to connect: %s", idx, err)
		return
	}

	defer conn.Close()
	buf := make([]byte, 1024)
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("client %d send %d", idx, i)
		l.Printf("%s\n", s)

		conn.Write([]byte(s))
		n, _ := conn.Read(buf)

		l.Printf("%s\n", string(buf[:n]))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	l := log.New(os.Stdout, "client: ", log.Ltime|log.Lmicroseconds)

	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	addr := os.Args[1]
	num, _ := strconv.Atoi(os.Args[2])
	wg := &sync.WaitGroup{}
	for i := 0; i < num; i++ {
		go start(l, i, wg, addr)
	}

	wg.Add(num)
	wg.Wait()
}
