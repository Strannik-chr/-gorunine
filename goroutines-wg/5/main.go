package main

import (
	"fmt"
	"sync"
	"time"
)

func barista(name string, drinkTime time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s начал готовить напиток\n", name)
	time.Sleep(drinkTime)
	fmt.Printf("%s закончил готовить напиток\n", name)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go barista("Ахьмад", 300*time.Millisecond,&wg)
	go barista("Зубайр", 500*time.Millisecond,&wg)
	go barista("Расул",200 * time.Millisecond,&wg)
	
	// Запусти трёх барист:
	// - Ахмад: 300ms
	// - Зубайр: 500ms
	// - Расул: 200ms

	wg.Wait()
	fmt.Println("Все напитки готовы!")
}
