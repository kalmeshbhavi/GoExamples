package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	fmt.Printf("Student %#v", student{
		Name: "John",
		Age:  10,
	})
}
