package main

import "fmt"

func main()  {
	ch := make(chan string,5)


ch <- "махито"
ch <- "миринда"
ch <- "торнато"
ch <- "тархун"
ch <- "компот"


fmt.Println(<- ch)
fmt.Println(<- ch)
fmt.Println(<- ch)
fmt.Println(<- ch)
fmt.Println(<- ch)
}