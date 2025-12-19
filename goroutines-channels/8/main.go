package main

import "fmt"

func sendNumbers(ch chan<- int) { // только отправка
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func printNumbers(ch <-chan int) { // только получение
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	ch := make(chan int)

	go sendNumbers(ch)
	printNumbers(ch)
}
