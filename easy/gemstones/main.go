// https://www.hackerrank.com/challenges/gem-stones/problem?isFullScreen=true
package main

import "fmt"

func gemstones(arr []string) int32 {
	// initialize map for rune and the frequency.
	m := make(map[rune]int, 0)

	// prepare output variable.
	var output int32

	// iterates array.
	for _, s := range arr {
		// create tmp map for avoid
		// multiple increment value.
		tmp := make(map[rune]struct{})

		// iterates string
		for _, r := range s {
			// if char already executed, skip it.
			if _, ok := tmp[r]; ok {
				continue
			}

			// increment char value.
			m[r] += 1
			// pushed executed char to tmp map.
			tmp[r] = struct{}{}
		}
	}

	// check if the value same with array length,
	// it's mean appeared on all array.
	for _, v := range m {
		if v == len(arr) {
			output++
		}
	}

	return output
}

func main() {
	fmt.Println(gemstones([]string{"abcdde", "baccd", "eeabg"}))
}
