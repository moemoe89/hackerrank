// https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree/problem?isFullScreen=true
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data}
	}
	if data <= root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}
	return root
}

func height(root *Node) int {
	if root == nil {
		return -1
	}
	leftHeight := height(root.left)
	rightHeight := height(root.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func main() {
	// Read input from stdin
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	numNodes, _ := strconv.Atoi(strings.TrimSpace(line1))
	line2, _ := reader.ReadString('\n')
	nums := strings.Split(strings.TrimSpace(line2), " ")
	var root *Node
	for i := 0; i < numNodes; i++ {
		num, _ := strconv.Atoi(nums[i])
		root = insert(root, num)
	}

	// Calculate and print height of binary tree
	fmt.Println(height(root))
}
