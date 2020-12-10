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
			_, exists := myctr[word]
			if exists {
				myctr[word] += 1
			} else {
				myctr[word] = 1
			}
		}

		for word, ctr := range myctr {
			if ctr > 1 {
				fmt.Printf("%s ocurred %d times\n", word, ctr)
			} else {
				fmt.Println(word, "ocurred", ctr, "time")
			}
		}
}
