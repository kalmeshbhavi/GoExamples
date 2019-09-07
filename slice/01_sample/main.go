package main

import "fmt"

func main() {
	var buffer [256]byte
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	addOneToEachElement(slice)
	fmt.Println("after", slice)
	copy := slice[1:5]
	copy[1] = 10
	fmt.Println("after copy", slice)
	fmt.Println("copy", copy)
	fmt.Println("len copy", len(copy))
	fmt.Println("copy[0]", copy[0])
	fmt.Println("slice[0]", slice[0])

}

func addOneToEachElement(bytes []byte) {
	for i := range bytes {
		bytes[i]++
	}
}
