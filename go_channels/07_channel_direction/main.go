package main

import "fmt"

func main() {
	c := incrementer()
	cSum := puller(c)

	for sum := range cSum {
		fmt.Println(sum)
	}
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func incrementer() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
