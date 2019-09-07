package main

import "fmt"

type test struct {
	id int
}

func main() {
	// t := new(test)
	arr := make([]test, 3)
	for i := 0; i < 3; i++ {
		t := &test{}
		t.id = i
		arr[i] = *t
	}

	fmt.Println(arr)
}
