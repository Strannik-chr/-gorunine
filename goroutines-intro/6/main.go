package main

import (
	"fmt"
	"time"
)


func main()  {
	go func ()  {
		for i := 5; i >= 1; i-- {
			fmt.Println(i,"...")
			time.Sleep(500*time.Millisecond)
		}
	}()

	time.Sleep(2*time.Second)
	fmt.Println("Поехали!")
}