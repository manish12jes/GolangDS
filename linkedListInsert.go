package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func traverseLinkList(head *Node) {
	for head != nil {
		fmt.Println("Element in list : ", head.data)
		head = head.next
	}
}

//Insert at begining of list
func insertAtBeginging(head *Node, value int) *Node {
	ptr := Node{}
	ptr.data = value
	ptr.next = head
	return &ptr
}

//Insert at index of list
func insertAtIndex(head *Node, index, value int) *Node {
	ptr := Node{}
	ptr.data = value
	p := head
	for i := 0; i != index-1; i++ {
		p = p.next
	}
	ptr.next = p.next
	p.next = &ptr

	return head
}

//Insert at end of list
func insertAtEnd(head *Node, value int) *Node {
	ptr := Node{}
	ptr.data = value
	p := head
	for p.next != nil {
		p = p.next
	}
	p.next = &ptr
	ptr.next = nil
	return head
}

//Insert after a specific node in list
func insertAfterNode(head, node *Node, value int) *Node {
	ptr := Node{}
	ptr.data = value
	ptr.next = node.next
	node.next = &ptr

	return head
}

func main() {
	var first, second, third, forth Node

	first.data = 1
	first.next = &second

	second.data = 2
	second.next = &third

	third.data = 3
	third.next = &forth

	forth.data = 4
	forth.next = nil
	head := &first

	fmt.Println("Before insert linked list")
	traverseLinkList(head)

	// head = insertAtBeginging(head, 0)
	head = insertAtIndex(head, 2, 8)
	// head = insertAtEnd(head, 5)
	// head = insertAfterNode(head, &second, 6)
	fmt.Println("\nAfter insert linked list")
	traverseLinkList(head)
}
