package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name     string
	TaskTime time.Duration
}

func submitTask(s Student, done chan<- string) {
	fmt.Printf("%s начал решать задачу\n", s.Name)
	time.Sleep(s.TaskTime)
	done <- fmt.Sprintf("%s сдал задачу!", s.Name)
}

func main() {
	students := []Student{
		{"Висхьаж", 300 * time.Millisecond},
		{"Iусман", 150 * time.Millisecond},
		{"Яхья", 400 * time.Millisecond},
		{"Амир", 200 * time.Millisecond},
		{"Шадид", 250 * time.Millisecond},
	}

	done := make(chan string)

	// TODO: запусти всех студентов
	for _, s := range students {
		go submitTask(s,done)
	}

	for i := 0; i < len(students); i++ {
		fmt.Println(<-done)
	}
	
	// TODO: получи len(students) сообщений из канала done

	fmt.Println("Все студенты сдали задания!")
}
