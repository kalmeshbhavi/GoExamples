package main

import "fmt"

func main() {
	n := 9
	dp := make([]int, n+1)
	fmt.Println("nth fib is : ", fib(n, dp))
}
func fib(n int, dp []int) int {
	if n <= 1 {
		dp[n] = n
	} else {
		dp[n] = fib(n-1, dp) + fib(n-2, dp)
	}
	return dp[n]
}
