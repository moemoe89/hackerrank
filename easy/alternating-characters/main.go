package main

import "fmt"

func alternatingCharacters(s string) int32 {
	var output int32

	for i := range s {
		// go to next iteration if i is first iteration.
		// because the first character doesn't have previous character.
		if i == 0 {
			continue
		}

		// if the previous character same with current character,
		// count the deletion.
		if s[i-1] == s[i] {
			output++
		}
	}

	return output
}

func main() {
	fmt.Println(alternatingCharacters("ABABABAB"))
}
