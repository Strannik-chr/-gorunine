package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
		wg.Add(1) 
		go func(n int) {
            defer wg.Done()
            fmt.Printf("Горутина %d\n", n)
        }(i)
    }

    wg.Wait()
    fmt.Println("Готово!")
}
//wg.Add(1) вызывается до запуска горутины, потому что иначе wg.Wait() может выполниться раньше увеличения счётчика и main завершится, не дождавшись горутины.