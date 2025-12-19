package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "один"
    }()

    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "два"
    }()

    // Ждём первый готовый канал
    select {
    case msg := <-ch1:
        fmt.Println("Получено из ch1:", msg)
    case msg := <-ch2:
        fmt.Println("Получено из ch2:", msg)
    }
}