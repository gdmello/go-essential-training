package main


import (
	"fmt"
	"net/http"
)

func returnType(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return "error"
	}

	defer resp.Body.Close()
	contentType := resp.Header.Get("content-type")
	return contentType
}

func main() {
	urls := []string {
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			ch <- returnType(url)
		}(url)

		ctype := <-ch
		fmt.Printf("%s -> %s\n", url, ctype)
	}
}