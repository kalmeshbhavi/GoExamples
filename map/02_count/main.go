package main

import "fmt"

func main() {
	count := frequencyCount([]string{"s", "a", "a", "b", "s"})
	fmt.Println(count)
}

func frequencyCount(skuIds []string) map[string]int {
	counts := make(map[string]int)
	for _, skuId := range skuIds {
		_, ok := counts[skuId]
		if ok {
			counts[skuId] += 1
		} else {
			counts[skuId] = 1
		}
	}
	return counts
}
