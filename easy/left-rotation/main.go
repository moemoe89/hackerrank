package main

import "fmt"

func rotateLeft(d int32, arr []int32) []int32 {
	n := int32(len(arr))

	d = d % n

	return append(arr[d:], arr[:d]...)
}

func main() {
	fmt.Println(rotateLeft(2, []int32{1, 2, 3, 4, 5}))
}
