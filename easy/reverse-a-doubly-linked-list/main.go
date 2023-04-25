// https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list/problem?isFullScreen=true
package main

import "fmt"

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

func reverse(llist *DoublyLinkedListNode) *DoublyLinkedListNode {
	if llist == nil {
		return llist
	}

	for llist != nil {
		prev := llist.prev
		llist.prev = llist.next
		llist.next = prev

		if llist.prev == nil {
			return llist
		}

		llist = llist.prev
	}

	return llist
}

func main() {
	first := &DoublyLinkedListNode{
		data: 1,
	}

	second := &DoublyLinkedListNode{
		data: 2,
	}

	third := &DoublyLinkedListNode{
		data: 3,
	}

	fourth := &DoublyLinkedListNode{
		data: 4,
	}

	first.next = second

	second.prev = first
	second.next = third

	third.prev = second
	third.next = fourth

	fourth.prev = third

	res := reverse(first)

	for res != nil {
		fmt.Printf("%+v\n", res)
		res = res.next
	}
}
