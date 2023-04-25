// https://www.hackerrank.com/challenges/reduced-string/problem?isFullScreen=true
package main

import (
	"fmt"
)

func superReducedString(s string) string {
	// initialize stack using array of rune.
	var stack []rune

	// iterates the string.
	for _, r := range s {
		// count the current stack length.
		n := len(stack)

		// if stack not empty, check last stack rune equal with current rune or not.
		if n > 0 && stack[n-1] == r {
			// if equal, it's mean the character should be deleted.
			stack = stack[:n-1]
		} else {
			// if empty, just add the rune to stack.
			stack = append(stack, r)
		}

		// example case 'aaabccddd'
		// [a] -> first stack input a
		// [] -> stack empty because 2nd char is a, remove last stack because it's duplicated
		// [a] -> third char pushed to stack
		// [a b] -> fourth char pushed to stack
		// [a b c] -> fifth char pushed to stack
		// [a b] -> c deleted because sixth char equal with last stack
		// [a b d] -> seventh char pushed to stack
		// [a b] -> d deleted becayse seventh char equal with last stack
		// [a b d] -> eight char pushed to stack
	}

	// if stack empty, return "Empty String"
	if len(stack) == 0 {
		return "Empty String"
	}

	// return the final string.
	return string(stack)
}

func main() {
	fmt.Println(superReducedString("aaabccddd"))
}
