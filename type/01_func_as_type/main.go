package main

type math func(i, j int) int

func main() {

}

func add(i, j int, math2 math) {
	math2(i, j)
}
