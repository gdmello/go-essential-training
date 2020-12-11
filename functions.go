package main

import (
	"fmt"
	"net/http"
	"strings"
)


func getContentType (url string) (string, error) {
	resp, error := http.Get(url)
	defer resp.Body.Close()
	if (error == nil) {
		for name, headers := range resp.Header {
			// for _, header := range headers {
				if (strings.ToLower(name) == strings.ToLower("Content-Type")) {
					return headers[0], error
				}
			// }
		}
	}
	return "", error
}

func main() {
	contentType, error := getContentType("https://golang.org")
	if (error == nil) {
		fmt.Println(contentType)
	}
}