package main

import "fmt"


func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
	// TODO: закрой канал после отправки всех чисел
}

func main() {
	ch := make(chan int)

	go sendNumbers(ch)

	for i := range ch{
	 
			fmt.Println(i)
		
	}
	// TODO: читай числа из канала, пока он не закрыт
	// Используй: for n := range ch { ... }
}
