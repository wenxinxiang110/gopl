package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopl/ch1"
	"gopl/ch12"
	"gopl/ch3"
	"gopl/ch4"
	"gopl/ch5"
	"gopl/ch6"
	"gopl/ch7"
	"gopl/ch8"
	"gopl/ch9"
)

func main() {
	chapter1()
	Server()
	// 读取输入到params这个slice中
	//params := make([]string, 0)
	//
	//in := bufio.NewScanner(os.Stdin)
	//
	//for in.Scan() {
	//
	//	params = append(params, in.Text())
	//}
	//fmt.Println(params)

	//Chapter4()

	//Chapter5(params[0])
	//Chapter6()
	//Server()

	//for _, item := range ch5.TopoSort(ch5.Preeqs) {
	//	fmt.Println(item)
	//}
	//Chapter7()
	//Chapter8()
	//Chapter9()
	//Chapter12()
}

func Server() {
	http.HandleFunc("/gif", ch1.GifHandler)

	http.HandleFunc("/corner", ch3.CornerHandler)
	http.HandleFunc("/complex", ch3.ComplexHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func chapter1() {
	if err := ch1.EchoOSArgs(os.Stdout, "\n"); err != nil {
		panic(err)
	}
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

		sucFlag = true

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

func Chapter6() {
	// 填充数据
	var x, y ch6.IntSet

	x.AddAll(1, 144, 9, 256)
	fmt.Println("set x:", x.String()) // "{1 9 144}"

	y.AddAll(9, 42, 144)
	fmt.Println("set y:", y.String()) // "{9 42}"

	fmt.Printf("len(x):%v ,len(y):%v\n", x.Len(), y.Len())

	fmt.Println("-------------------------------------------")

	// 并集
	u := x.Copy()
	u.UnionWith(&y)
	fmt.Println("x union y:", u.String()) // "{1 9 42 144}"

	fmt.Println("-------------------------------------------")

	// 交集
	i := x.Copy()
	i.IntersectWith(&y)
	fmt.Println("x intersect y:", i.String()) // "{9}"
	i2 := y.Copy()
	i2.IntersectWith(&x)
	fmt.Println("y intersect x:", i2.String()) // "{9}"

	fmt.Println("-------------------------------------------")

	d := x.Copy()
	d.DifferenceWith(&y)
	fmt.Println("x difference y:", d.String()) // "{1 144}"
	d2 := y.Copy()
	d2.DifferenceWith(&x)
	fmt.Println("y difference x:", d2.String()) // "{1 144}"

	fmt.Println("-------------------------------------------")

	s := x.Copy()
	s.SymmetricDifference(&y)
	fmt.Println("x symmetric difference y:", s.String()) // "{1 42 144}"

}

func Chapter7() {

	//time.Sleep(time.Second)
	//ch7.PrintTracks(ch7.Tracks)

	db := ch7.Database{Datas: map[string]ch7.Dollars{"shoes": 50, "socks": 5}}

	// 自定义路由选择器 ,这段有点脱裤子放屁
	//mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.List))
	//mux.Handle("/price", http.HandlerFunc(db.Price))
	//
	//log.Fatal(http.ListenAndServe("localhost:8000", mux))

	// 等价于这段代码
	//mux.HandleFunc("/list", db.List)
	//mux.HandleFunc("price", db.Price)
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func Chapter8() {

	//ch8.ClockServerAndClient()
	//ch8.Cal()
	ch8.PipDemo()
}

func Chapter9() {
	ch9.Cal()
}

func Chapter12() {
	ch12.Type()
}
