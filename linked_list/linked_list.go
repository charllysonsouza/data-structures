package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(value interface{}) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

func (l *LinkedList) Remove(value interface{}) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
	}

	currentNode := l.head
	for currentNode.next != nil {
		if currentNode.next.value == value {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

func (l *LinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%v->", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Print("nil")
}
