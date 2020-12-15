package main

import (
	"fmt"
	"io"
	"bytes"
	"strings"
)

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	s:= string(p)
	s = strings.ToUpper(s)
	newbytes := []byte(s)
	c.wtr.Write(newbytes)
	return len(newbytes), nil
}

func main() {
	var buf bytes.Buffer
	capper := &Capper{wtr: &buf}
	bytelen, err := capper.Write([]byte("this is all miXEd cAsE writing 1"))
	if err == nil {
		fmt.Println(buf.String())
		fmt.Println(bytelen)
	} else {
		panic(err)
	}
	fmt.Println(byte('A' - 'a'))
}