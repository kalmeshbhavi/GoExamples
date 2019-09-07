package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T ", s[i])
	}

	fmt.Println(strings.Count(s, "l"))

	fmt.Println("\n", s)
}
