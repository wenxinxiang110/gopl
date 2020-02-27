// 一個可以向服务端读写的简单客户端程序
package cli

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan struct{})

	// 起一个后台goroutines读取服务端的内容到标准输出
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	// 读取标准输入的内容，写入到conn中
	if _, err := io.Copy(conn, os.Stdin); err != nil {
		log.Fatal(err)
	}

	<-done
}
