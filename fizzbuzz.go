package main

import (
    "fmt"
)

func main() {
   for i:=1; i<21;i++ {
	switch {
		case  ((i%3) == 0) && ((i%5) == 0):
		case  (i%3) == 0:
			fmt.Println("fizz (3)")
			fmt.Println("fizz (3)")
		case (i%5) == 0:
			fmt.Println("buzz (5)")
	}
	   fmt.Println(i)
   }
}
