package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("BMW лучший!")
	}()
	go func() {
		fmt.Println("MERS лучший!")
	}()
	go func() {
		fmt.Println("AUDI лучший!")
	}()

	time.Sleep(time.Second)

}
