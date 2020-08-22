package ch7

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Dollars float32

func (d Dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// 一个模拟的简易数据库
type Database struct {
	Datas map[string]Dollars
	sync.RWMutex
}

// 继承Handler接口
func (d Database) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/list":
		d.List(w, req)
	case "/price":
		d.Price(w, req)
	case "/create":
		d.Create(w, req)
	default:
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound)
	}

}

func (d Database) List(w http.ResponseWriter, req *http.Request) {
	d.RLock()
	defer d.Unlock()
	for item, price := range d.Datas {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d Database) Price(w http.ResponseWriter, req *http.Request) {
	d.RLock()
	defer d.Unlock()

	item := req.URL.Query().Get("item")

	price, ok := d.Datas[item]

	if !ok {
		msg := fmt.Sprintf("no such item:%q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func (d Database) Create(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("item")
	value := req.URL.Query().Get("price")

	price, err := strconv.ParseFloat(value, 32)

	if err != nil {
		http.Error(w, fmt.Sprintf("parse price failed:%v", err), http.StatusOK)
		return
	}

	d.Lock()
	defer d.Unlock()

	d.Datas[key] = Dollars(price)

	fmt.Fprintf(w, "create success,[%s]=%v", key, price)

}
