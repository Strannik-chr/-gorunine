package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumber(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина #%d\n", n)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		time.Sleep(100 * time.Millisecond)
		go printNumber(i, &wg)
	}

	wg.Wait()
	fmt.Println("Все горутины завершились!")
}
