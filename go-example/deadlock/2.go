package main

func main() {
    var ch chan int = make(chan int, 2)
    ch <- 1
    ch <-2
    //ch <-3  会报deadlock
}
