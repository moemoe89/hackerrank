package main

import "fmt"

func plusMinus(arr []int32) {
	p := 0
	n := 0
	z := 0

	for i := range arr {
		if arr[i] > 0 {
			p++
		} else if arr[i] == 0 {
			z++
		} else {
			n++
		}
	}

	t := float32(len(arr))

	fmt.Println(float32(p) / t)
	fmt.Println(float32(n) / t)
	fmt.Println(float32(z) / t)
}

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}
