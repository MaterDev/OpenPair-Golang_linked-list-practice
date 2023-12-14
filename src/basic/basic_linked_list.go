package basic

import "fmt"

type Node struct {
	data int // Primary data
	next *Node // A address to the next node in the list.
}

type LinkedList struct {
	head *Node
}

// Create new node at the end of the list
func (list *LinkedList) AddToEnd(data int) {
	// ! 1. Make a new node to add to the list
	newNode := &Node{data: data}

	// ! 2. If the list is empty, then create the head.
	if list.head == nil {
		list.head = newNode
		return
	}

	last := list.head

	// ! 3. Traversal loop - will assign last to be .next until it reaches the end of the list.
	for last.next != nil {
		last = last.next
	}
	// ! 4 . Add new address to the node that was created at the beginning of the method.
	last.next = newNode
}

// Loop over list and display all data
func (list *LinkedList) Display() {
	// Starting Point
	current := list.head

	for current != nil {
		fmt.Println(current.data, " ")
		current = current.next
	} 
}

func BasicLinkedList() {
	fmt.Println("Basic Linked List!")

	// Initialize list
	list := LinkedList{}
	list.AddToEnd(1)
	list.AddToEnd(2)
	list.AddToEnd(3)
	list.AddToEnd(4)
	list.AddToEnd(5)
	list.AddToEnd(6)
	list.Display() // will output all data => 1 2 3 4 5 6
}