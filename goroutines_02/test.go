package main

import (
	"fmt"
)

type etag *string

type test struct {
	etag etag
}

func main() {
	str := "test"
	etag := &str
	t := &test{etag: etag}
	fmt.Println(*t.etag)
}
