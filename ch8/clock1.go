package ch8

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const addr = "localhost:8000"

func ClockServerAndClient() {
	go Server(ClockHandleConn)
	go Client()

	time.Sleep(time.Minute)

}

func Server(f func(net.Conn)) {
	// 可以用telnet监听， telnet localhost 8000
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server start")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go f(conn)
	}
}

func ClockHandleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}

func Client() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Client start")

	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
