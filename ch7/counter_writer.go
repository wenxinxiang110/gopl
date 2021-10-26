package ch7

import (
	"io"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

type CountingWriterFactory interface {
	CountingWriter(w io.Writer) (io.Writer, *int64)
}

type CountingWriterFactoryImpl struct{}

func (c *CountingWriterFactoryImpl) CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CountingWriter{
		Writer: w,
	}
	return cw, &cw.count
}

type CountingWriter struct {
	io.Writer
	count int64
}

func (c *CountingWriter) Write(p []byte) (n int, err error) {
	c.count += int64(len(p))
	return c.Writer.Write(p)
}
