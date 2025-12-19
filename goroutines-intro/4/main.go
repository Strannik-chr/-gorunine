package main

import (
    "fmt"
    "time"
)

func main() {
    names := []string{"Ахмад", "Зубайр", "Расул", "Асхаб", "Человек паук"}

    for _, name := range names {
        go func() {
            fmt.Println("Привет,", name)
        }()
    }

    time.Sleep(time.Second)
}