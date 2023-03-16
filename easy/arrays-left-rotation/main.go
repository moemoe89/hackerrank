package main

import "fmt"

func rotLeft(a []int32, d int32) []int32 {
	return append(a[d:], a[:d]...)
}

func main() {
	fmt.Println(rotLeft([]int32{1, 2, 3, 4, 5}, 4))
}
