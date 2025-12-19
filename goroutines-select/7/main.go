package main

import (
	"fmt"
	"time"
)

func player(name string, table chan int, done <-chan struct{}) {
	for {
		select {
		case ball := <-table:
			fmt.Printf("%s получил мяч #%d\n", name, ball)
			time.Sleep(100 * time.Millisecond)

			if ball >= 10 {
				return
			}

			table <- ball + 1

		case <-done:
			fmt.Printf("%s уходит с корта\n", name)
			return
		}
	}
}

func main() {
	table := make(chan int)
	done := make(chan struct{})

	go player("Имран", table, done)
	go player("Зубайр", table, done)

	table <- 1 // запускаем мяч

	for {
		ball := <-table
		if ball >= 10 {
			close(done)
			break
		}
		table <- ball
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Игра окончена!")
}
