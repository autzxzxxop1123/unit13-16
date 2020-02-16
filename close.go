package main

import "fmt"

func say(ch chan string) {
	ch <- "Hello"
	ch <- "Hola"
	ch <- "Bonjour"
	ch <- "Dia dhuit"
	ch <- "Ciao"
	close(ch)
}

func main(){
	ch := make(chan string, 5)
	go say(ch)
	for i := range ch {
		fmt.Println(i)
	}
}