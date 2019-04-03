package main

import "fmt"

func main() {
	fmt.Println(getName("test1"))
}

func getName(str string) string {

	for   {
		fmt.Println("running")
		if str == "test" {
			return "success"
		}

		if str == "test1" {
			return "success1"
		}

	}
}
