package main

import (
    "fmt"
)

func main() {
	book := "MyBible Old and New testament"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("value: %v type %T\n", book, book)
	fmt.Printf("value: %v type %T\n", book[0], book[0])
	fmt.Printf("value: %v type %T\n", book[1], book[1])

	fmt.Println(book[:1])
	fmt.Println(book[len(book)-1:len(book)])
	sbook := fmt.Sprintf("%d", 1001)
	fmt.Println(sbook[0])
	fmt.Println(sbook[len(sbook)-1])
}
