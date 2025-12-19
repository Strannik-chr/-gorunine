package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		time.Sleep(500 * time.Millisecond) // слишком долго!
// 		ch <- "результат"
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Получено:", msg)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("Таймаут! Слишком долго ждали.")
// 	}
// }

func fetchWithTimeout(ch <-chan string, timeout time.Duration) (string, bool) {
	ch = make(chan string)

	select {
	case msg := <-ch:
		return msg, true
	case <-time.After(timeout):
		return "", false
	}

}

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "данные"
	}()

	if result, ok := fetchWithTimeout(ch, 1*time.Second); ok {
		fmt.Println("Успех:", result)
	} else {
		fmt.Println("Таймаут!")
	}
}
