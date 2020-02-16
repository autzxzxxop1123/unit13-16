package main

import "fmt"

func sum(c chan int, number ...int) {
	sum := 0 
	for _, v := range number {
		sum = sum + v
	}
	c <- sum
}

func printer(c1, c2 chan int) {
	select {
	case num1 := <-c1:
		fmt.Println("channe1-1 :", num1)
	case num2 := <-c2:
		fmt.Println("channe1-2 :", num2)
	}
}

func main(){ 
	c1 := make(chan int)
	c2 := make(chan int)
	go printer(c1, c2)
	//go sum(c1, 1,2,3)
	go sum(c2, 10,11)

	var input string
	fmt.Scan(&input)
}