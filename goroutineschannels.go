package main


import (
	"fmt"
	"net/http"
)

func returnType(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("error:%s", err)
		return 
	}

	defer resp.Body.Close()
	contentType := resp.Header.Get("content-type")
	ch <- fmt.Sprintf("%s -> %s", url, contentType)
}

func main() {
	urls := []string {
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch) //use go keyword to generate a goroutine
	}

	for range urls {
		ctype := <-ch
		fmt.Printf(ctype)
	}
}