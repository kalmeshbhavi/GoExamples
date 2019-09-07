package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	c := boring("joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	for i := 0; ; i++ {
		select {
		case c <- fmt.Sprintf("%s: %d", msg, i):
		// Do nothing
		case <-quit:
			return c
		}
	}
}
