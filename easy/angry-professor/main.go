package main

import "fmt"

func angryProfessor(k int32, a []int32) string {
	n := int32(len(a))

	var numLate int32
	for _, v := range a {
		if v > 0 {
			numLate++
		}
	}

	attend := n - numLate
	if attend < k {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Println(angryProfessor(3, []int32{-2, -1, 0, 1, 2}))
}
