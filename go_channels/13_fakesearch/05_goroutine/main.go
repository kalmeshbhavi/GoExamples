package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

type Result string

var (
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2")
	Image1 = fakeSearch("Image1")
	Image2 = fakeSearch("Image2")
	Video1 = fakeSearch("Video1")
	Video2 = fakeSearch("Video2")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
