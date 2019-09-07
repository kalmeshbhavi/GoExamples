package main

import "fmt"

func main() {

	var t test
	t = &test1{}
	t.test()

	t = &test2{}
	t.test()
}

type test interface {
	test()
}

type test1 struct {
}

func (t test1) test() {
	fmt.Println("test1")
}

type test2 struct {
}

func (t test2) test() {
	fmt.Println("test2")
}
