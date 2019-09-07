package main

import "fmt"

func main() {
	fmt.Println(noOfWays([]int{1, 4, 2, 1, 5}))
}

func noOfWays(arr []int) int {
	options := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			if options != 0 {
				return 0
			}
			if i == 1 || arr[i-2] <= arr[i] {
				options++
			}
			if i == len(arr)-1 || arr[i-1] <= arr[i+1] {
				options++
			}
			if options == 0 {
				return 0
			}
		}
	}
	if options == 0 {
		return len(arr)
	}
	return options
}
