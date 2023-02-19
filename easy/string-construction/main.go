package main

import "fmt"

func stringConstruction(s string) int32 {
	mapS := make(map[rune]int)

	var out int32

	for _, v := range s {
		if _, ok := mapS[v]; !ok {
			mapS[v] = 1
			out++
		}
	}

	return out
}

func main() {
	fmt.Println(stringConstruction("abcd"))
}
