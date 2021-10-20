package ch1

import (
	"fmt"
	"io"
	"os"
)

// EchoOSArgs 输出命令行参数
func EchoOSArgs(w io.Writer, sep string) error {
	for idx, arg := range os.Args[1:] {
		if _, err := w.Write([]byte(fmt.Sprintf("args %d: %s%s", idx+1, arg, sep))); err != nil {
			return err
		}
	}
	return nil
}
