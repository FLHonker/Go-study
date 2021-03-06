package main

import "fmt"

func Count(ch chan int) {
    ch <- 1
    fmt.Println("counting.")
}

func main() {
    chs := make([]chan int, 10)
    for i := 0; i < 10; i++ {
        chs[i] = make(chan int)
        go Count(chs[i])
    }

    for _, ch := range(chs) {
        <-ch
    }
	/* 随机0,1
    ch := make(chan int, 1)
    for {
        select {
        case ch <- 0:
        case ch <- 1:
        }
        i := <-ch
        fmt.Println("Value recv:", i)
    }
    */
}
