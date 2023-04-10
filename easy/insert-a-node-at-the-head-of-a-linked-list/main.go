// https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list/problem?isFullScreen=true
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

func insertNodeAtHead(llist *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{
		data: data,
		next: llist,
	}
}

func printLinkedList(llist *SinglyLinkedListNode) {
	for llist != nil {
		fmt.Println(llist.data)

		llist = llist.next
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line1))

	var llist *SinglyLinkedListNode
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		data, _ := strconv.Atoi(strings.TrimSpace(line))
		llist = insertNodeAtHead(llist, int32(data))
	}
	printLinkedList(llist)
}
