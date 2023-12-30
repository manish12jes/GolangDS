package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func traverseLinkList(head *Node) {
	for head.next != nil {
		fmt.Println("Element in list : ", head.data)
		head = head.next
	}
}

// Delete 1st node from the list
func deleteFirstNode(head *Node) *Node {
	p := head
	head = p.next
	p = nil
	return head
}

// Delete node at specific index from the list
func deleteAtIndex(head *Node, index int) *Node {
	p := head
	for i := 0; i < index-1; i++{
		p = p.next
	}
	p.next = p.next.next
	p.next.next = nil
	return head
}

func deleteLastNode(head *Node) *Node {
	p := head
	q := &Node{}
	for p.next != nil {
		q = p
		p = p.next
	}
	q.next = nil
	p = nil
	return head
}

func deleteNodeByValue(head *Node, value int) *Node {
	p := head
	q := &Node{}
	for p.data != value {
		q = p
		p = p.next
	}

	q.next = p.next
	p = nil
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

	fmt.Println("Before insert link list")
	traverseLinkList(head)

	head = deleteFirstNode(head)
	// head = deleteAtIndex(head, 2)
	// head = deleteLastNode(head)
	// head = deleteNodeByValue(head, 3)
	fmt.Println("\nAfter delete link list")
	traverseLinkList(head)
}