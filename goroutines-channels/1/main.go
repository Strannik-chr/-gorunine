package main

import "fmt"

func main() {
ch := make(chan string)

go func ()  {
	ch <- "привет из горутин!"
}()

go func ()  {
	ch <- "привет из intocodes!"
	
}()


msg := <- ch
msg2 := <- ch
fmt.Println(msg,msg2)
}