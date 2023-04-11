// https://www.hackerrank.com/challenges/insert-a-node-into-a-sorted-doubly-linked-list/problem?isFullScreen=true
package main

import "fmt"

// Definition for doubly-linked list:
type DoublyLinkedListNode struct {
	data int32
	prev *DoublyLinkedListNode
	next *DoublyLinkedListNode
}

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	newNode := &DoublyLinkedListNode{data: data}

	// Special case for an empty list
	if llist == nil {
		return newNode
	}

	// Special case for inserting at the beginning of the list
	if data < llist.data {
		newNode.next = llist
		llist.prev = newNode
		return newNode
	}

	curr := llist
	for curr.next != nil && curr.next.data < data {
		curr = curr.next
	}

	// Insert node after the current node
	newNode.next = curr.next
	if curr.next != nil {
		curr.next.prev = newNode
	}
	curr.next = newNode
	newNode.prev = curr

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

	sortedInsert(first, 5)

	for first != nil {
		fmt.Printf("%+v\n", first)
		first = first.next
	}
}
