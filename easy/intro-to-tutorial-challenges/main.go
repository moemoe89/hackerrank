package main

import "fmt"

func introTutorial(V int32, arr []int32) int32 {
	var out int32

	for k, v := range arr {
		if v == V {
			out = int32(k)

			break
		}
	}

	return out
}

func main() {
	fmt.Println(introTutorial(4, []int32{1, 4, 5, 6, 9, 12}))
}
