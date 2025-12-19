package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("привет из горутины!")
}

func main() {
	go sayHello()
	fmt.Println("привет из main!")
	time.Sleep(time.Second)
}