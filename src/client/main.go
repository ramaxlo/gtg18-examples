package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var fail_count int32

func fail() {
	atomic.AddInt32(&fail_count, 1)
}

func usage() {
	fmt.Println("Usage: ./client <addr> <num of connections>")
}

func start(l *log.Logger, idx int, wg *sync.WaitGroup, addr string) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fail()
		}
	}()

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

		n, err := conn.Write([]byte(s))
		if err != nil {
			l.Printf("client %d fail to write: %s\n", idx, err)
			panic("write fail")
		}

		n, err = conn.Read(buf)
		if err != nil {
			l.Printf("client %d fail to read: %s\n", idx, err)
			panic("read fail")
		}

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
		time.Sleep(1 * time.Millisecond)
		go start(l, i, wg, addr)
		wg.Add(1)
	}

	wg.Wait()
	l.Printf("Fail count: %d\n", fail_count)
}
