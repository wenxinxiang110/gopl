package ch8

import (
	"github.com/adonovan/gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

//下面的程序会循环迭代一些图片文件名，并为每一张图片生成一个缩略图：
func Make(filenames []string) {

	// 这里的ch作用相当于sync.WaitGroup
	ch := make(chan struct{})

	for _, f := range filenames {
		go func(file string) {
			thumbnail.ImageFile(file)
			ch <- struct{}{}
		}(f)

	}

	for range filenames {
		<-ch
	}
}

func Make2(filenames <-chan string) int64 {
	sizes := make(chan int64)

	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()

		}(f)

	}
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
