// Go file
/***********************************************
# Copyright (c) 2018, Wuhan
# All rights reserved.
#
# @Filename: gdbtest.go
# @Version：V1.0
# @Author: Frank Liu - frankliu624@gmail.com
# @Description: ---
# @Create Time: 2018-07-05 14:34:24
# @Last Modified: 2018-07-05 14:34:24
************************************************/

package main

import (
	"fmt"
	"time"
)

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}

func main() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gofunc"
	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}
