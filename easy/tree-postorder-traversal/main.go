// https://www.hackerrank.com/challenges/tree-postorder-traversal/problem?isFullScreen=true
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func postorderTraversal(root *Node) {
	if root == nil {
		return
	}

	postorderTraversal(root.left)
	postorderTraversal(root.right)
	fmt.Print(root.data, " ")
}

func insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data}
	} else {
		var cur *Node
		if data <= root.data {
			cur = insert(root.left, data)
			root.left = cur
		} else {
			cur = insert(root.right, data)
			root.right = cur
		}
		return root
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var n int
	_, err := fmt.Fscan(reader, &n)
	if err != nil {
		panic(err)
	}

	var root *Node
	for i := 0; i < n; i++ {
		var data int
		_, err := fmt.Fscan(reader, &data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		root = insert(root, data)
	}

	postorderTraversal(root)
}
