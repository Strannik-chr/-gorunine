package main

import "fmt"

func sum(numbers []int, ch chan<- int) {
	total := 0
	for _, n := range numbers {
		total += n
	}
	ch <- total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int)

	// Раздели срез пополам и запусти две горутины
	go sum(numbers[:5], ch) // сумма первой половины
	go sum(numbers[5:], ch) // сумма второй половины

	// TODO: получи два результата и сложи их
	// Ожидаемый результат: 55

	n := <-ch
	h := <-ch
	fmt.Println(n + h)

}
