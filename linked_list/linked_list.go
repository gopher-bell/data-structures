package main

import (
	"fmt"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	next *Node
	prev *Node
	data int
}

func (l *LinkedList) addFront(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head.prev = n
		l.head = n
	}

	l.length++
	return
}

func (l *LinkedList) deleteFront() {
	if l.head == nil {
		return
	} else {
		l.head.next.prev = nil
		l.head = l.head.next
	}

	l.length--
	return
}

func (l *LinkedList) deleteBack() {
	if l.tail == nil {
		return
	} else {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	}

	l.length--
	return
}

func (l *LinkedList) delete(key int) {
	if l.length == 0 {
		return
	}
	if l.head.data == key {
		l.deleteFront()
		return
	}
	if l.tail.data == key {
		l.deleteBack()
		return
	}

	prev := l.head
	for prev.next.data != key {
		if prev.next.next == nil {
			return
		}
		prev = prev.next
	}

	prev.next = prev.next.next
	l.length--
	return
}

func (l *LinkedList) addBack(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}

	l.length++
	return
}

func (l LinkedList) print() {
	n := l.head
	for n != nil {
		fmt.Print(n.data, "-> ")
		n = n.next
	}
}

func main() {
	l := &LinkedList{}

	n1 := &Node{data: 5}
	n2 := &Node{data: 4}
	n3 := &Node{data: 7}
	n4 := &Node{data: 100}
	n5 := &Node{data: 54}
	n6 := &Node{data: 71}
	n7 := &Node{data: 11}
	n8 := &Node{data: 22}
	n9 := &Node{data: 33}

	l.addFront(n1)
	l.addFront(n2)
	l.addFront(n3)
	l.addFront(n4)
	l.addFront(n5)
	l.addFront(n6)
	l.addBack(n7)
	l.addBack(n8)
	l.addBack(n9)

	l.print()
	fmt.Println()
	fmt.Println(l.length)
	fmt.Println(l.head.data)
	fmt.Println(l.tail.data)

	l.delete(333)

	l.print()
	fmt.Println()
	fmt.Println(l.length)
	fmt.Println(l.head.data)
	fmt.Println(l.tail.data)
}
