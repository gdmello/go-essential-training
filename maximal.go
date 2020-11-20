package main

import (
    "fmt"
)

func main() {
   nums :=  []int{16,8,42,4,23,15,43}
   max :=  nums[1]
   for  i :=0; i < len(nums) ;i++ {
	   if max < nums[i] {
		   max = nums[i]
	   } 
   }

   fmt.Println(max)
}
