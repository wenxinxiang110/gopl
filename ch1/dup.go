package ch1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// DupUtil 模仿uniq的命令行工具，可以支持多个文件处理
func DupUtil(out io.Writer) {
	var (
		args      []string
		filePaths []string
	)
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			args = append(args, arg)
		}
		filePaths = append(filePaths, arg)
	}

	for _, path := range filePaths {
		file, err := os.Open(path)
		if err != nil {
			panic(fmt.Sprintf("Open file %s error:%v", path, err))
		}
		_, _ = out.Write([]byte(path + "\n"))
		_ = Dup(file, out, DupOptions{args})
		_ = file.Close()
	}

}

// DupOptions todo: 扩展功能，后续用于支持其他参数
type DupOptions struct {
	OriginParams []string
}

// Dup 类试uniq的工具
func Dup(r io.Reader, w io.Writer, _ DupOptions) error {
	scanner := bufio.NewScanner(r)
	counter := make(map[string]int)
	if scanner.Scan() {
		counter[scanner.Text()]++
	}
	for line, count := range counter {
		_, _ = w.Write([]byte(fmt.Sprintf("%d %s\n", count, line)))
	}
	return nil
}
