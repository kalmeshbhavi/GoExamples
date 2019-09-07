package main

import "encoding/json"

type test struct {
	name string
}

func main() {

	// var t test
	bytes, _ := json.Marshal(nil)
	println(string(bytes))

}
