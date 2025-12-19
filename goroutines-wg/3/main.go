package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Воркер %d завершился\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i,&wg)
	}

	wg.Wait()
	fmt.Println("Готово!")
}

//sync.WaitGroup не копируют, не возвращают, не передают по значению — только по *sync.WaitGroup