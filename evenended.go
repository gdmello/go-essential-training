package main

import (
    "fmt"
)

func main() {
	ctr := 0
	for a:=10; a<=9999;a++ {
		for x:= a; x<=9999; x++ {
			y := a*x
			//fmt.Println(y)
			sy := fmt.Sprintf("%d", y)
			lsy := len(sy)
			//fmt.Println(sy)
			//fmt.Printf("Start: %v End: %v\n",sy[:1],sy[lsy-1:lsy] )
			if sy[0] == sy[lsy-1] {
			    ctr++
			}
		}
        }
	fmt.Println("Number of evenended numbers between 100 & 9999 %v", ctr)

}
