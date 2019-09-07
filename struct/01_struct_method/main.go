package main

import "fmt"

func main() {
	e := Employee{
		FirstName: "John",
		LastName:  "Smith",
	}

	e.PhoneNum = 123
	fmt.Println(e.fullName())
	e.changeName("ABC")
	fmt.Println(e.fullName())
	fmt.Println(e.getNum())
	fmt.Println(e.Contact)
}

type Contact struct {
	PhoneNum int
}

type Employee struct {
	FirstName string
	LastName  string
	Contact
}

func (e *Employee) fullName() string {
	return fmt.Sprintf("%s %s", e.FirstName, e.LastName)
}

func (e *Employee) changeName(lastName string) {
	e.LastName = lastName
}

func (e *Employee) getNum() int {
	return e.PhoneNum
}
