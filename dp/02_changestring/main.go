package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(noOfStepsToChangeString("Kalmesh", "Kasturi"))
}

func noOfStepsToChangeString(str1 string, str2 string) int {
	secStrLen := len(str2)
	firstStrLen := len(str1)
	rowLen := len(str1) + 1
	colLen := len(str2) + 1

	dp := make([][]int, rowLen)
	for i := range dp {
		dp[i] = make([]int, colLen)
	}

	for row := 0; row <= len(str2); row++ {
		dp[row][0] = row
	}

	for col := 0; col <= len(str1); col++ {
		dp[0][col] = col
	}

	for row := 1; row <= secStrLen; row++ {
		for col := 1; col <= firstStrLen; col++ {
			ind := 0
			if str1[col-1] != str2[row-1] {
				ind = 1
			}
			dp[row][col] = min(dp[row][col-1]+1, dp[row-1][col]+1, dp[row-1][col-1]+ind)
		}
	}
	for row := range dp {
		fmt.Println(dp[row])
	}

	return dp[len(str2)][len(str1)]
}

func min(arr ...int) int {
	sort.Ints(arr)
	return arr[0]
}
