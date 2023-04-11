// https://www.hackerrank.com/challenges/merge-two-sorted-linked-lists/problem?isFullScreen=true
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

func mergeLists(llist1 *SinglyLinkedListNode, llist2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	llists := []int32{}

	for llist1 != nil || llist2 != nil {

		if llist1 != nil {
			llists = append(llists, llist1.data)

			llist1 = llist1.next
		}

		if llist2 != nil {
			llists = append(llists, llist2.data)

			llist2 = llist2.next
		}
	}

	quickSort(llists, 0, len(llists)-1)

	var merged *SinglyLinkedListNode

	for i := range llists {
		merged = insertNodeAtTail(merged, llists[i])
	}

	return merged
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

func quickSort(arr []int32, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []int32, left int, right int) int {
	pivot := arr[right]

	i := left

	for j := left; j < right; j++ {
		// < operator means ascending sort
		// > operator means descending sort
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
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
		merged := mergeLists(llist1, llist2)

		// Print the result
		for merged != nil {
			fmt.Print(merged.data, " ")
			merged = merged.next
		}

		fmt.Println()
	}
}
