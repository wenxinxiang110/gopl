package main

import (
	"errors"
	"gopl/ch1"
	"gopl/ch3"
	"gopl/ch4"
	"log"
	"net/http"
	"os"
)

func main() {
	CH4()

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

func CH4() {
	var result *ch4.IssuesSearchRes
	var err = errors.New("init")
	for err != nil {
		result, err = ch4.SearchIssues(ch4.IssuesParam)
		if err != nil {
			log.Printf("failed:%v", err)
			continue
		}
		// 打印到控制台
		ch4.PrintIssues(result, os.Stdout)

		// 写入文件
		WriteIssuesIntoFile(result)
	}
}

func WriteIssuesIntoFile(is *ch4.IssuesSearchRes) {
	file, e := os.OpenFile("issues.html", os.O_CREATE|os.O_WRONLY, os.ModeDir)
	if e != nil {
		log.Printf("failed to open file issues.html:%v", e)
		return
	}
	ch4.PrintIssuesHtml(is, file)
	file.Close()

}
