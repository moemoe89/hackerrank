// https://www.hackerrank.com/challenges/tree-inorder-traversal/problem?isFullScreen=true
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

func inorderTraversal(root *Node) {
	if root == nil {
		return
	}

	inorderTraversal(root.left)
	fmt.Print(root.data, " ")
	inorderTraversal(root.right)
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

	inorderTraversal(root)
}
