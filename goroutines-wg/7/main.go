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
		fmt.Println("Работаю...")
		// забыли wg.Done()
		//происходит ошибака  exit status 2/зависание 
	}()

	wg.Wait()
	fmt.Println("Готово!")
}
