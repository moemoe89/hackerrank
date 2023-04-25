// https://www.hackerrank.com/challenges/permutation-equation/problem?isFullScreen=true
package main

import "fmt"

func permutationEquation(p []int32) []int32 {
	var out []int32

	mapP := make(map[int32]int32)
	for k, v := range p {
		mapP[v] = int32(k + 1)
	}

	for k, _ := range p {
		k = k + 1
		out = append(out, mapP[mapP[int32(k)]])
	}

	return out
}

func main() {
	fmt.Println(permutationEquation([]int32{2, 3, 1}))
}
