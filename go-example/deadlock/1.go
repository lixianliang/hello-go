package main

func main() {
    ch := make(chan int)
    //<- ch
    ch <- 1
}
