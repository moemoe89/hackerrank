// https://www.hackerrank.com/challenges/game-of-thrones/problem?isFullScreen=true
package main

import "fmt"

func gameOfThrones(s string) string {
	// initialize the map.
	m := make(map[rune]int, len(s))

	// store all element and the frequency to map.
	for _, r := range s {
		m[r]++
	}

	// prepare the odd num variable.
	oddNum := 0

	// iterates the map.
	for _, v := range m {
		// checks the frequency odd or not.
		// if odd increment the value.
		if v%2 != 0 {
			oddNum++

			// if odd number found more than 1,
			// it definitely not palindrome.
			if oddNum > 1 {
				return "NO"
			}
		}
	}

	return "YES"
}

func main() {
	fmt.Println(gameOfThrones("aaabbbb"))
}
