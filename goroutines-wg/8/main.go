package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer wg.Done() // лишний Done!
		fmt.Println("Работаю...")
	}()

	wg.Wait()
	fmt.Println("Готово!")
}
//Каждый Done уменьшает этот счётчик на 1.
//Если  вызовать Done() два раза, счётчик уйдёт в минус 
//go заметить и вызавит панику 