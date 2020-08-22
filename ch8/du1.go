package ch8

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// sema is a counting semaphore for limiting concurrency in dirents.
// 就是用来限制协程数量，不要无限制暴涨
var sema = make(chan struct{}, 20)

var done = make(chan struct{})

// 递归/多线程搜索文件目录
func Run(roots []string) {

	// 读取输入就关闭done的协程
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// ...determine roots...
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()
	// ...select loop...
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()

	if canceller() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {

	select {
	case sema <- struct{}{}:
	case <-done:
		return nil

	}
	// 获取和释放

	defer func() {
		<-sema
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}

	return entries
}

// 要增加一个退出的功能
func canceller() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
