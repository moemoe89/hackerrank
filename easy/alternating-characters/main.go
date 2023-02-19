package main

import "fmt"

func alternatingCharacters(s string) int32 {
	var output int32

	beforeChar := ""

	for i, r := range s {
		if i > 0 {
			beforeChar = string(s[i-1])
		}

		if len(beforeChar) > 0 && beforeChar == string(r) {
			output++

			continue
		}
	}

	return output
}

func main() {
	fmt.Println(alternatingCharacters("ABABABAB"))
}
