package main

import (
	"fmt"
	"sync"
	"time"
)

type Student struct {
	Name     string
	TaskTime time.Duration
}

func submitTask(s Student, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s начал решать задачу\n", s.Name)
	time.Sleep(s.TaskTime)
	fmt.Printf("%s сдал задачу!\n", s.Name)
}

func main() {
	students := []Student{
		{"Висхьаж", 300 * time.Millisecond},
		{"Iусман", 150 * time.Millisecond},
		{"Яхья", 400 * time.Millisecond},
		{"Амир", 200 * time.Millisecond},
		{"Шадид", 250 * time.Millisecond},
	}

	var wg sync.WaitGroup
	
	for _,stud := range students{
		wg.Add(1)
		go submitTask(stud,&wg)
	}
	wg.Wait()
	fmt.Println("Все студенты сдали задания!")
}
