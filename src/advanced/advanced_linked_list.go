package advanced

import "fmt"

type Node struct {
	data int // Primary data
	next *Node // A address to the next node in the list.
}

type LinkedList struct {
	head *Node
}

// Add a new node at the end of the linked list
func (llist *LinkedList) AddToEnd(data int) {
	newNode := &Node{data: data}

	if llist.head == nil {
		llist.head = newNode
		return
	}
	
	last := llist.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode

}

func AdvancedLinkedList() {
	fmt.Println("Advanced Linked List!")
}