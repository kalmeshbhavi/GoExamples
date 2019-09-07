package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("kalmesh"), boring("Kasturi"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("you both are boring, I am leaving")
}

func fanIn(input1 <-chan string, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(s string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}
	}()
	return out
}
