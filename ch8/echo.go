package ch8

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func Echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprint(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, strings.ToLower(shout))
}

func EchoHandleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		Echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func EchoClient() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
