package main

import "fmt"

func main() {
	fmt.Println("nth fib : ", fib(2))
}

func fib(n int) int {
	fmt.Println(n)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
