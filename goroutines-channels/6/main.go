package main

func main() {
    ch := make(chan int)
    close(ch)
    close(ch)
}

// panic: close of closed channel — возникает, когда код пытается закрыть канал повторно (канал можно закрывать только один раз).

// канал можно закрывать только один раз