package main

import (
	"fmt"
	"time"
)

func printMessage(msg string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("%s: %d\n", msg, i)
		time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	go printMessage("медленная", 3)
	go printMessage("быстрая", 10)

	time.Sleep(11 * time.Second)
}
// нам предется годать сколько сек поставить 
//хотя в этом задании я интуитивно понял что 11 