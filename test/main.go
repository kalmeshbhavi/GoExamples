package main

import "fmt"

type intTest interface {
	Hello()
}

type test1 struct {
}

type test2 struct {
}

func (t1 test1) Hello() {
	fmt.Println("test1")
}
func (t2 test2) Hello() {
	fmt.Println("test2")
}

func print(intT intTest) {
	intT.Hello()
}

type test3 struct{}

func main() {
	print(test1{})
}

func getName(str string) string {

	for {
		fmt.Println("running")
		if str == "test" {
			return "success"
		}

		if str == "test1" {
			return "success1"
		}

	}
}
