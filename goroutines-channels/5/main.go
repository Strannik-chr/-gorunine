package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "первое"
	ch <- "второе"
	close(ch)

// for ok := <- ch {
// 	if  ok == false{
// 	fmt.Println("канал закрыт")
// 	} else {
// 		fmt.Println("канал открыт")
// 	}

	// TODO: прочитай все значения, используя цикл for и проверку ok

	// Когда ok == false, выведи "Канал закрыт!" и выйди из цикла
	value, ok := <-ch
	

	if !ok {
		fmt.Println("канал закрыт и пуст")
	}else{
		fmt.Println("значение получено")
	}
	fmt.Println(value)
		
	
}
