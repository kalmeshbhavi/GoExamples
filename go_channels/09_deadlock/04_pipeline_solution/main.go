package main

import "fmt"

func main() {
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}

func sq(ins chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range ins {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}
