package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func printLinkedList(head *SinglyLinkedListNode) {
	// loop until the end of linked list.
	for {
		// print the element of each linked list.
		fmt.Println(head.data)

		// if reach the last element, break the loop.
		if head.next == nil {
			break
		}

		// override the current (head) with the next.
		head = head.next
	}
}

func main() {
	llist := &SinglyLinkedListNode{
		data: 1,
		next: &SinglyLinkedListNode{
			data: 2,
			next: &SinglyLinkedListNode{
				data: 3,
			},
		},
	}

	printLinkedList(llist)
}
