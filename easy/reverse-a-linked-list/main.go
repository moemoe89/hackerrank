// https://www.hackerrank.com/challenges/reverse-a-linked-list/problem?isFullScreen=true
package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	var prev, next *SinglyLinkedListNode

	curr := llist

	for curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next
	}

	return prev
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

	llist = reverse(llist)

	// Print the elements of the linked list
	for llist != nil {
		fmt.Printf("%+v\n", llist)
		llist = llist.next
	}
}
