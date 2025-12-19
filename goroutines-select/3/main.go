package main

import "fmt"

func main() {
    ch := make(chan string)

    select {
    case msg := <-ch:
        fmt.Println("Получено:", msg)
    default:
        fmt.Println("Нет данных")
    }

    ch1 := make(chan string,1)


	ch1 <- "привет"
	 select {
    case msg := <-ch1:
        fmt.Println("Получено:", msg)
    default:
        fmt.Println("Нет данных")
    }
}