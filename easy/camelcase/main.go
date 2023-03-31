package main

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int32 {
	output := int32(1)

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			output++
		}
	}

	return output
}

func camelcase2(s string) int32 {
	output := int32(1)

	for _, r := range s {
		if unicode.IsUpper(r) {
			output++
		}
	}

	return output
}

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
}
