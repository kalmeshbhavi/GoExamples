package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	var list LinkedList
	fmt.Println("empty list : ", list.IsEmpty())
	list.Append(1)
	list.Append(2)
	list.Append(4)
	err := list.Insert(3, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("list elements...")
	list.String()
	fmt.Println("size : ", list.size)
	fmt.Println("Head : ", list.Head().content)
	fmt.Println("Index of 4 : ", list.IndexOf(4))
	i, err := list.RemoveAt(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Removed value : ", i.content)
	list.String()
	list.reverse()
	fmt.Println("After reverse")
	list.String()

}
