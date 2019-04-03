package main

import "fmt"

func main() {

	fmt.Println("Execution started")
	go printHello()
	fmt.Println("Execution Completed")
}

func printHello() {
	fmt.Println("Hello")
}