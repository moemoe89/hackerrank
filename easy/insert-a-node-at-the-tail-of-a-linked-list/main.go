// https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list/problem?isFullScreen=true
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeAtTail(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	node := &SinglyLinkedListNode{
		data: data,
	}

	if head == nil {
		return node
	}

	curr := head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = node

	return head
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read the number of elements in the linked list
	n, err := strconv.Atoi(readLine(reader))
	if err != nil {
		panic(err)
	}

	var head *SinglyLinkedListNode

	// Read the elements of the linked list
	for i := 0; i < n; i++ {
		data, err := strconv.Atoi(readLine(reader))
		if err != nil {
			panic(err)
		}
		head = insertNodeAtTail(head, int32(data))
	}

	// Print the elements of the linked list
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(str))
}
