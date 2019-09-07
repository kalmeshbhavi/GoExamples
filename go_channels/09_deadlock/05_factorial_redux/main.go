package main

import "fmt"

func main() {
	f := factorial(24)
	for v := range f {
		fmt.Print(v)
	}
}

func factorial(n int) chan int {

	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
