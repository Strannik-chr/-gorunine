package main

import (
	"fmt"
	"time"
)

func barista(name string, drinkTime time.Duration) {
	fmt.Printf("%s начал готовить напиток\n", name)
	time.Sleep(drinkTime)
	fmt.Printf("%s закончил готовить напиток\n", name)
}

func main() {
	go barista("Ахьмад", 300*time.Microsecond)
	go barista("Зубайр", 500*time.Microsecond)
	go barista("Расул", 200*time.Microsecond)

time.Sleep(1 * time.Second)

}
