package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
