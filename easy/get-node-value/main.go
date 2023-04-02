// https://www.hackerrank.com/challenges/get-the-value-of-the-node-at-a-specific-position-from-the-tail/problem?isFullScreen=true
package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func getNode(llist *SinglyLinkedListNode, positionFromTail int32) int32 {
	var values []int32

	// loop until the end of linked list.
	for {
		values = append(values, llist.data)

		// if reach the last element, break the loop.
		if llist.next == nil {
			n := int32(len(values))

			// then return the value based on position from tail.
			return values[n-1-positionFromTail]
		}

		// override the current (llist) with the next.
		llist = llist.next
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

	fmt.Println(getNode(llist, 2))
}
