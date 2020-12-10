package main

import (
    "fmt"
    "strings"
)

func main() {
		text := `Needles and pins
		needles and pins
		Sew me a sail
		To catch me the wind`

        var myctr = make(map[string]int)
		for _, word  := range strings.Fields(text) {
			myctr[strings.ToLower(word)]++
		}
		fmt.Println(myctr)
}
