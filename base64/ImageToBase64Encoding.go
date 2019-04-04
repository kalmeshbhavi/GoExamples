package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("./base64/test.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buf := make([]byte, size)
	fReader := bufio.NewReader(file)
	_, _ = fReader.Read(buf)
	base64Str := base64.StdEncoding.EncodeToString(buf)

	fmt.Println(base64Str)

}
