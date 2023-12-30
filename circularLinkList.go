package main

import "fmt"

//
type Node struct {
	data	int
	next	*Node
}


func traverseLinkList(head *Node) {
	p := head
	for {
		fmt.Println("Element in list : ", p.data)
		p = p.next
		if p == head {
			break
		}
	}
}

func insertAtBeginging(head *Node, value int) *Node {
	p := head
	ptr := Node{}
	for p.next != head{
		p = p.next
	}
	p.next = &ptr
	ptr.data = value
	ptr.next = head
	
	return &ptr
}

func insertAtIndex(head *Node, index, value int) *Node {
	p := head
	ptr := Node{}
	for i := 0; i < index-1; i++ {
		p = p.next
	}
	ptr.data = value
	ptr.next = p.next
	p.next = &ptr
	return head
}

func insertAtEnd(head *Node, value int) *Node{
	p := head
	for {
		p = p.next
		if p.next == head {
			break
		}
	}
	ptr := Node{}
	ptr.data = value
	p.next = &ptr
	ptr.next = head
	return head
}

func insertAfterNode(head, second *Node, value int) *Node{
	ptr := Node{}
	ptr.data = value
	ptr.next = second.next
	second.next = &ptr
	return head
}

func main() {
	var first, second, third, forth  Node
	first.data = 1
	first.next = &second

	second.data = 2
	second.next = &third

	third.data = 3
	third.next = &forth
	
	forth.data = 4
	forth.next = &first
	head := &first

	fmt.Println("Before insert linked list")
	traverseLinkList(head)

	// head = insertAtBeginging(head, 5)
	// head = insertAtIndex(head, 2, 6)
	// head = insertAtEnd(head, 7)
	head = insertAfterNode(head, &second, 8)
	fmt.Println("\nAfter insert linked list")
	traverseLinkList(head)
}