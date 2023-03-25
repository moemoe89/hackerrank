package main

import "fmt"

func utopianTree(n int32) int32 {
	var h int32

	for i := int32(0); i <= n; i++ {
		if i%2 == 0 {
			h = h + 1
		} else {
			h = h * 2
		}
	}

	return h
}

func main() {
	fmt.Println(utopianTree(5))
}
