package main

import (
	"bufio"
	"bytes"
	"gopl/ch1"
	"gopl/ch3"
	"gopl/ch4"
	"gopl/ch5"
	"log"
	"net/http"
	"os"
)

func main() {

	params := make([]string, 0)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {

		params = append(params, in.Text())
	}
	//fmt.Println(params)

	//Chapter4()

	Chapter5(params[0])

	//Server()

	//for _, item := range ch5.TopoSort(ch5.Preeqs) {
	//	fmt.Println(item)
	//}

}

func Server() {
	http.HandleFunc("/gif", ch1.GifHandler)

	http.HandleFunc("/corner", ch3.CornerHandler)
	http.HandleFunc("/complex", ch3.ComplexHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func Chapter4() {

	// 是否获取成功
	var sucFlag bool

	for !sucFlag {
		result, err := ch4.SearchIssues(ch4.IssuesParam)
		if err != nil {
			log.Printf("failed:%v", err)
			continue
		}
		// 打印到控制台
		ch4.PrintIssues(result, os.Stdout)

		// 写入文件
		file, e := os.OpenFile("issues.html", os.O_CREATE|os.O_WRONLY, os.ModeDir)
		defer file.Close()
		if e != nil {
			log.Printf("failed to open file issues.html:%v", e)
			return
		}
		ch4.PrintIssuesHtml(result, file)

	}
}

func Chapter5(url string) {

	var buf bytes.Buffer

	ch1.Fetch(url, &buf)

	ch5.FindLinks(&buf)

}
