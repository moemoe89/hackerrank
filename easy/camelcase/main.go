// https://www.hackerrank.com/challenges/camelcase/problem?isFullScreen=true
package main

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int32 {
	// Because the first word is lowercase, then initialize 1.
	output := int32(1)

	for _, r := range s {
		// Checks if the rune is uppercase (from A to Z).
		if r >= 'A' && r <= 'Z' {
			// increment the word since the next word
			// starting with uppercase.
			output++
		}
	}

	return output
}

func camelcase2(s string) int32 {
	// Because the first word is lowercase, then initialize 1.
	output := int32(1)

	for _, r := range s {
		// Checks if the rune is uppercase.
		if unicode.IsUpper(r) {
			// increment the word since the next word
			// starting with uppercase.
			output++
		}
	}

	return output
}

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
}
