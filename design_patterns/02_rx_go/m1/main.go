package main

import "time"

func main() {
	Method2()
}

func Method2() {
	items := make(chan interface{})
	go loadItems(items)

	for data := range items {
		if d, ok := data.(int); ok {
			if d < 10 {
				go func() {
					process(d)
				}()
			} else {
				go func() {
					process(d * 10)
				}()
			}
		}
	}
	// s1 := src.Filter(isInt).Filter(ltTen).Subscribe(ob)

}

func process(i int) {
	println(i)
}

func loadItems(items chan interface{}) {
	for i := 0; i < 5; {
		i++
		time.Sleep(time.Second)
		items <- i
	}
	close(items)
}
