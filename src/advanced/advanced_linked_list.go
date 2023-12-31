package advanced

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

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

func (llist *LinkedList) Display() {
	current := llist.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println()
}

// Replace the head with a new node
func (llist *LinkedList) AddToStart(data int) {
	//! Create new node, with .next pointing to the original head
	newNode := &Node{data: data, next: llist.head}
	//! Replace head with the new node
	llist.head = newNode
}

// Delete 1st node from list based on matching with data
func (llist *LinkedList) DeleteWithValue(data int) {
	// If the list is empty, stop
	if llist.head == nil {
		return
	}

	if llist.head.data == data {
		llist.head = llist.head.next
		return
	}

	current := llist.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

//! Reverse linked list
func (llist *LinkedList) Reverse() {
	var prev, next *Node
	current := llist.head

	for current != nil {
		next = current.next
		current.next = prev
		fmt.Printf("current.next: %v \n", current.next)
		fmt.Printf("prev: %v \n", prev)
		prev = current
		current = next
	}

	llist.head = prev
}

// TODO: Create a method that will delete all nodes with value arg

func AdvancedLinkedList() {
	fmt.Println("Advanced Linked List!")

	list := LinkedList{}
	list.AddToEnd(1)
	list.AddToEnd(2)
	list.AddToEnd(3)
	list.AddToEnd(0)
	list.Display() // Output: 1, 2, 3, 0

	list.AddToStart(0)
	list.Display() // Output: 0, 1, 2, 3, 0

	// ! Will only delete either the head + 1 node, or first node that has the value
	list.DeleteWithValue(0)
	list.Display() // Output:  1, 2, 3, 0

	list.Reverse()
	list.Display() // Output: 0, 3, 2, 1
}
