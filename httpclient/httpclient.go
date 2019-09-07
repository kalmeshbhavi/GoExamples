package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	req := []byte(`{
		"path" : "tmp_photo/photo/oZ8yUDVqUJhtXeXNcgkYic2nj9VrHGiPJMhoTsSRQF3m.jpg"
	}`)

	response, err := http.Post("http://localhost:1112/services/photo/v1/download", "", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println("err ", err)
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("err 1 ", err)
	}

	fmt.Println("response ", bytes)
}
