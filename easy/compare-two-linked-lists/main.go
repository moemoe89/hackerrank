// https://www.hackerrank.com/challenges/compare-two-linked-lists/problem?isFullScreen=true
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Definition for singly-linked list.
type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func compareLists(llist1 *SinglyLinkedListNode, llist2 *SinglyLinkedListNode) int32 {
	// If both lists are nil, they are equal
	if llist1 == nil && llist2 == nil {
		return 1
	}

	// If only one list is nil, they are not equal
	if llist1 == nil || llist2 == nil {
		return 0
	}

	// If the current elements are not equal, the lists are not equal
	if llist1.data != llist2.data {
		return 0
	}

	// Recursively compare the rest of the lists
	return compareLists(llist1.next, llist2.next)
}

// Fill in this function to compare two linked lists
func compareLists2(llist1 *SinglyLinkedListNode, llist2 *SinglyLinkedListNode) int32 {
	for llist1 != nil || llist2 != nil {
		if llist1.data != llist2.data {
			return 0
		}

		if llist1.next != nil && llist2.next == nil {
			return 0
		}

		if llist2.next != nil && llist1.next == nil {
			return 0
		}

		llist1 = llist1.next
		llist2 = llist2.next
	}

	return 1
}

func main() {
	// Read the input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		var llist1 *SinglyLinkedListNode
		for j := 0; j < n; j++ {
			scanner.Scan()
			data, _ := strconv.Atoi(scanner.Text())
			llist1 = &SinglyLinkedListNode{data: int32(data), next: llist1}
		}
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		var llist2 *SinglyLinkedListNode
		for j := 0; j < m; j++ {
			scanner.Scan()
			data, _ := strconv.Atoi(scanner.Text())
			llist2 = &SinglyLinkedListNode{data: int32(data), next: llist2}
		}

		// Compare the two linked lists
		result := compareLists(llist1, llist2)

		// Print the result
		fmt.Println(result)
	}
}
