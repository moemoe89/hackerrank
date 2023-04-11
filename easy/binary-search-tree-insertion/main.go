// https://www.hackerrank.com/challenges/binary-search-tree-insertion/problem?isFullScreen=true
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

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	dataStr := strings.Split(scanner.Text(), " ")
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i], _ = strconv.Atoi(dataStr[i])
	}

	var root *Node
	for i := 0; i < n; i++ {
		root = insert(root, data[i])
	}

	preOrder(root)
}
