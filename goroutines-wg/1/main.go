package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Привет из горутины!")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello(&wg)

	fmt.Println("Привет из main!")

	wg.Wait()
}
