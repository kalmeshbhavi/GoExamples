package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Execution started")
	go printHello()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("Execution Completed")
}

func printHello() {
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Hello")
}