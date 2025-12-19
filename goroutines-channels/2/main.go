package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		ch <- 42
	}()

	msg := <- ch
	fmt.Println(msg)
}
