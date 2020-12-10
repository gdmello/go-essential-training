package main

import (
    "fmt"
    "strings"
)

func main() {
		text := `Needles and pins
		Needles and pins
		Sew me a sail
		To catch me the wind`

        var myctr = make(map[string]int)
		for _, word  := range strings.Fields(text) {
			myctr[word]++
		}
		fmt.Println(myctr)
}
