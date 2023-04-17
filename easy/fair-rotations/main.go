// https://www.hackerrank.com/challenges/fair-rations/problem?isFullScreen=true
package main

import "fmt"

func fairRations(B []int32) string {
	n := len(B) - 1

	count := int32(0)

	for i := 0; i < n; i++ {
		if B[i]%2 == 0 {
			continue
		}

		B[i]++
		B[i+1]++

		count += 2
	}

	if B[n]%2 == 1 {
		return "NO"
	}

	return numbersToStr(count)
}

func numbersToStr(n int32) string {
	if n == 0 {
		return "0"
	}

	// prepare the array of rune
	var str []rune

	// iterates until numbers become 0
	for n != 0 {
		// use modulo 10 in order to get every digit
		// e.g. 12345 % 10 -> 5
		m := n % 10

		// convert to rune string
		r := '0' + rune(m)

		// use prepend because we need to add
		// digit from right to left
		// [5] -> [4 5] -> [3 4 5]
		str = append([]rune{r}, str...)

		// divide the number to eliminate the last digit
		// 12345 to 1234
		n /= 10
	}

	return string(str)
}

func main() {
	fmt.Println(fairRations([]int32{2, 3, 4, 5, 6}))
}
