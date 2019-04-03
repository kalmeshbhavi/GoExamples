package main

import (
	"fmt"
	"time"
)

func main()  {
	dateTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(dateTime)

	x := (30/100) * (1066000*(4/15) + 94341)
	fmt.Println(x)
}
