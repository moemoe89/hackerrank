// https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list-in-reverse/problem?isFullScreen=true
package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func reversePrint2(llist *SinglyLinkedListNode) {
	var values []int32

	// loop until the end of linked list.
	for {
		// append to array with reverse position,
		values = append([]int32{llist.data}, values...)

		// if reach the last element, break the loop.
		if llist.next == nil {
			break
		}

		// override the current (llist) with the next.
		llist = llist.next
	}

	for i := range values {
		fmt.Println(values[i])
	}
}

func reversePrint(llist *SinglyLinkedListNode) {
	if llist == nil {
		return
	}

	reversePrint(llist.next)

	fmt.Println(llist.data)
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

	reversePrint(llist)
}
