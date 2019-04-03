package main

import "fmt"

func greet(c chan string) {
	fmt.Println(<- c)
	fmt.Println(<- c)
}

func main() {
	fmt.Println("main() started")

	c := make(chan string, 1)

	go greet(c)
	c <- "Kalmesh"

	close(c)
	c <- "test"

	fmt.Println("main() stopped")
}
