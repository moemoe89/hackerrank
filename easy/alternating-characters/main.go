package main

import "fmt"

func alternatingCharacters(s string) int32 {
	var output int32

	var beforeChar byte

	for i := range s {
		if i > 0 {
			beforeChar = s[i-1]
		}

		if beforeChar > 0 && beforeChar == s[i] {
			output++

			continue
		}
	}

	return output
}

func main() {
	fmt.Println(alternatingCharacters("ABABABAB"))
}
