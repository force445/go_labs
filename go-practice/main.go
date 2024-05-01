package main

import (
	"fmt"
	"go-practice/routine"
	"time"
)

func main() {
	book1 := book{
		bookID: "1556",
		Title:  "The Deep",
	}
	defer fmt.Println(print_book_name(book1))

	start := time.Now()
	ch := make(chan string)
	go routine.Routine("hello", 10, ch)
	go routine.Routine("bye", 10, ch)

	message := <-ch

	if message == "sending into channel" {
		fmt.Println("go routine success")
		fmt.Println("Time to execuse: ", time.Since(start), "second")
	}
}

func print_book_name(mybook book) string {
	return mybook.Title
}
