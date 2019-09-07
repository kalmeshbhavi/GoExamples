package main

import (
	"fmt"
	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
	"time"
)

type item string

func main() {
	Execute()
}

func Execute() {

	items := make(chan interface{})
	go loadItems(items)

	isInt := func(i interface{}) bool {
		_, ok := i.(int)
		return ok
	}
	ltTen := func(i interface{}) bool {
		return i.(int) < 10
	}
	gtTen := func(i interface{}) bool {
		return i.(int) > 10
	}
	timesTen := func(i interface{}) interface{} {
		return i.(int) * 10
	}

	it, _ := iterable.New(items)

	src := observable.From(it)

	process := handlers.NextFunc(func(item interface{}) {
		fmt.Printf("Processing: %v\n", item)
	})
	cleanup := handlers.DoneFunc(func() {
		println("Done")
	})

	ob := observer.New(process, cleanup)

	//
	s1 := src.Filter(isInt).Filter(ltTen).Subscribe(ob)
	s2 := src.Filter(isInt).Filter(gtTen).Map(timesTen).Subscribe(ob)

	select {
	case <-s1:
	case <-s2:
	}
}

func loadItems(items chan interface{}) {
	for i := 0; i < 15; {
		i++
		time.Sleep(time.Second)
		items <- i
	}
	close(items)
}
