package main

import "fmt"

func main() {

	g := gen()

	f := factorial(g)
	for v := range f {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func factorial(in <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for v := range in {
			out <- fact(v)
		}
		close(out)
	}()
	return out

}
func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
