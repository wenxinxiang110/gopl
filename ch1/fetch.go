package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
		return
	}

	// todo:attention this ,使用Copy方法直接输出到一个writter对象中(这里是一个标准输出)
	// 其他，例如使用原生的http包，也可以直接输出到一个http.ResponseWriter对象中
	_, err = io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url, err)
		return
	}

	//fmt.Println()
}
