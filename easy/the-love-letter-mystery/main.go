// https://www.hackerrank.com/challenges/the-love-letter-mystery/problem?isFullScreen=true
package main

import "fmt"

func theLoveLetterMystery(s string) int32 {
	// initialize left and right index.
	i := 0
	j := len(s) - 1

	// initialize the output reduce number.
	output := int32(0)

	// iterates until left and right index same.
	for i <= j {
		a := rune(s[i])
		b := rune(s[j])

		// another approach, calculates the reductions
		// by decrement the current rune and target rune.
		//if b > a {
		//	b, a = a, b
		//}
		//
		//for a != b {
		//	a--
		//	output++
		//}

		// get the diff for both rune, it can be negative if,
		// b > a, then need to make it absolute.
		reductions := int(a - b)
		if reductions < 0 {
			reductions = -reductions
		}

		// accumulates reduce number.
		output += int32(reductions)

		// sliding 1 position for left and right index.
		i++
		j--
	}

	return output
}

func main() {
	fmt.Println(theLoveLetterMystery("abc"))
}
