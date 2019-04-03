package main

import (
	null "gopkg.in/guregu/null.v3"
	"fmt"
	"strings"
	"time"
)

const dateTimeFormat = "2006-01-02 15:04:05"

func main()  {

	var colorIDS []uint64
	str := null.StringFrom(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(colorIDS)), ","), "[]"))
	fmt.Println(str.String)

	dateTime := time.Now().Format(dateTimeFormat)
	fmt.Println(dateTime)

	emptyTime := time.Time{}
	var t time.Time

	if emptyTime.IsZero() {
		fmt.Println("empty")
	} else {
		fmt.Println("failure")
	}

	if t == emptyTime {
		fmt.Println("empty")
	} else {
		fmt.Println("failure")
	}

	var ui uint64 = 18446744073709551615
	fmt.Print(ui)
}
