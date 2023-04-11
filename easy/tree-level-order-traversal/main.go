// https://www.hackerrank.com/challenges/tree-level-order-traversal/problem?isFullScreen=true
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

func levelOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		fmt.Print(curr.data, " ")

		if curr.left != nil {
			queue = append(queue, curr.left)
		}

		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
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

	levelOrderTraversal(root)
}
