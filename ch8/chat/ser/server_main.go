// 一个模拟聊天室，这是服务端
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	// 服务端地址
	ServerAddr = "localhost:8000"
)

type client chan<- string

var (
	// entering和leaving这两个队列代表客户端进出的动作，消费掉就没了
	entering = make(chan client)
	leaving  = make(chan client)

	// message里的消息会被copy出来广播每一个连接上的cli
	messages = make(chan string)
)

func main() {
	// 启动服务端
	listener, err := net.Listen("tcp", ServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	// 广播事件协程，专门处理各个客户端之间的状态
	go broadcaster()

	// 处理连接过来的每一个客户端
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// 处理连接
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	// outgoing client messages

	ch := make(chan string)

	go clientWriter(c, ch)

	// 客户端地址
	who := c.RemoteAddr().String()

	ch <- "You are " + who
	// 通知所有所有客户端我进来聊天室了
	messages <- who + " has arrived"

	entering <- ch

	// 读取客户端的输入到messages协程
	input := bufio.NewScanner(c)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	// 当客户端结束输入的时候，向所有客户端广播退出的消息
	leaving <- ch

	messages <- who + " has left"

	// 释放连接
	c.Close()
	close(ch)
}

// 把ch中的内容写入客户端中
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcaster() {

	online := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range online {
				cli <- msg
			}
		case cli := <-entering:
			online[cli] = true

		case cli := <-leaving:
			delete(online, cli)
			//close(cli)
		}
	}
}
