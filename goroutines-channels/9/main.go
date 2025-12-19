package main

import (
    "fmt"
    "time"
)

func barista(name string, orders <-chan string, done chan<- string) {
    for order := range orders {
        fmt.Printf("%s готовит %s...\n", name, order)
        time.Sleep(300 * time.Millisecond)
        done <- fmt.Sprintf("%s готов (сделал %s)", order, name)
    }
}

func main() {
    orders := make(chan string, 5)
    done := make(chan string, 5)

    // Запусти двух барист
    go barista("Ахмад", orders, done)
    go barista("Зубайр", orders, done)

    // Отправь 5 заказов
    drinks := []string{"Латте", "Капучино", "Эспрессо", "Американо", "Шур тоьх"}
    for _, drink := range drinks {
        orders <- drink
    }
    close(orders) // заказов больше не будет

	for i := 0; i < len(drinks);i++{
		fmt.Println(<- done)
	}
    // TODO: получи и выведи 5 готовых напитков из канала done
}