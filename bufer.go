package main

import "fmt"

func main() {

	ch := make(chan string, 2 )
	ch <- "Hello"
	ch <- "Hi"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}