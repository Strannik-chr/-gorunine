package main

import "fmt"

func trySend(ch chan<- int, value int) bool {
	ch = make(chan int)

	select {
	case ch <- value:
		return true
	default:
		return false

	}

	// TODO: используй select с default

	// Верни true, если удалось отправить, false — если канал занят
}

func main() {
	ch := make(chan int, 1)

	fmt.Println(trySend(ch, 1)) // true — буфер пуст
	fmt.Println(trySend(ch, 2)) // false — буфер полон

	<-ch                        // освобождаем буфер
	fmt.Println(trySend(ch, 3)) // true — снова есть место
}
