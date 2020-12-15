package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strconv"
)

func killServer(pidFile string) error {
	content, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "can't open PID file")
	}
	// defer content.Close()
	newStr := string(content)
	fmt.Println(newStr)
	iPid, _ := strconv.Atoi(newStr)
	fmt.Println("killing ", iPid)
	return nil
}

func main() {
	err := killServer("./pid.file")
	if err != nil {
		panic("ERROR")
	}
}