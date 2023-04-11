// https://www.hackerrank.com/challenges/delete-duplicate-value-nodes-from-a-sorted-linked-list/problem?isFullScreen=true
package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func removeDuplicates(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	curr := llist

	for curr.next != nil {
		if curr.data == curr.next.data {
			curr.next = curr.next.next

			continue
		}

		curr = curr.next
	}

	return llist
}

func printLinkedList(head *SinglyLinkedListNode) {
	// loop until the end of linked list.
	for head != nil {
		// print the element of each linked list.
		fmt.Println(head.data)

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
				next: &SinglyLinkedListNode{
					data: 3,
				},
			},
		},
	}

	removeDuplicates(llist)
	printLinkedList(llist)
}
