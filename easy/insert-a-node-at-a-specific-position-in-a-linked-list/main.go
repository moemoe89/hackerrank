// https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem?isFullScreen=true
package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	tmp := &SinglyLinkedListNode{
		data: data,
	}

	head := llist

	idx := int32(0)

	for idx < position-1 {
		head = head.next
		idx++
	}

	tmp.next = head.next
	head.next = tmp

	return llist
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

	res := insertNodeAtPosition(llist, 4, 2)

	for res != nil {
		fmt.Printf("%+v\n", res)
		res = res.next
	}
}
