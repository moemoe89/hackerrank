package main

import "fmt"

func makeAnagram(a string, b string) int32 {
	mapAB := make(map[rune]int32)

	for _, r := range a {
		mapAB[r]++
	}

	for _, r := range b {
		mapAB[r]--
	}

	output := int32(0)

	for _, val := range mapAB {

		if val < 0 {
			output += -(val)
		} else {
			output += val
		}

	}

	return output
}

func main() {
	fmt.Println(makeAnagram("cde", "abc"))
}
