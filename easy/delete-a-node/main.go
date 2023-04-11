// https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list/problem?isFullScreen=true
package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	curr := llist

	i := int32(-1)

	for curr.next != nil {
		i++

		if i == position {
			curr.data = curr.next.data
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

	deleteNode(llist, 2)
	printLinkedList(llist)
}
