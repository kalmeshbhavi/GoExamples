package linkedlist

import (
	"fmt"
	"sync"
)

type Node struct {
	content int
	next    *Node
}

type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (l *LinkedList) Append(data int) {
	l.lock.Lock()
	node := Node{data, nil}
	if l.head == nil {
		l.head = &node
	} else {
		last := l.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	l.size++
	l.lock.Unlock()
}

func (l *LinkedList) Insert(data int, i int) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if i < 0 || i > l.size {
		return fmt.Errorf("index out of bounds")
	}
	addNode := Node{data, nil}
	if i == 0 {
		addNode.next = l.head
		l.head = &addNode
		return nil
	}
	node := l.head
	j := 0
	for j < i-2 {
		j++
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode
	l.size++
	return nil
}

func (l *LinkedList) RemoveAt(i int) (*Node, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if i < 0 || i > l.size {
		return nil, fmt.Errorf("index out of bounds")
	}
	node := l.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	l.size--
	return remove, nil
}

func (l *LinkedList) IndexOf(i int) int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	node := l.head
	j := 0
	for {
		if node.content == i {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}

func (l *LinkedList) IsEmpty() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	if l.head == nil {
		return true
	}
	return false
}

func (l *LinkedList) Size() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	size := 1
	last := l.head
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}
	return size
}

func (l *LinkedList) Head() *Node {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.head
}

func (l *LinkedList) String() {
	l.lock.RLock()
	defer l.lock.RUnlock()
	node := l.head
	for {
		if node == nil {
			break
		}
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}

func (l *LinkedList) reverse() {
	l.lock.Lock()
	defer l.lock.Unlock()
	var prev *Node
	cur := l.head
	var next *Node
	for cur != nil {
		// Store next
		next = cur.next
		// Reverse current node's pointer
		cur.next = prev
		// Move pointers one position ahead.
		prev = cur
		cur = next
	}
	l.head = prev
}
