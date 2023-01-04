package main

import (
    "fmt"
)

// This Node represtents a particular node of the linkedlist
type Node struct {
	value int
	next  *Node
}

// LinkedList represents the entire linked list
type LinkedList struct {
	head *Node
	len  int
}


// Insert inserts new node at the end of the linked list
func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}


// GetAt returns node at given position of the linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}



// InsertAt inserts new node at given position.
//we have to traverse the list till the provided nth node & change it’s pointer to point to the new node
func (l *LinkedList) InsertAt(pos int, value int) {
	// create a new node
	newNode := Node{}
	newNode.value = value
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}


// Print displays all the nodes of the linked list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

func main(){
    link := LinkedList{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)

    
	fmt.Printf("Head: %v\n", link.head.value)
	
	link.Print()
	fmt.Println("\n==============================\n")

	link.InsertAt(2,1334)
	link.Print()

	
}
