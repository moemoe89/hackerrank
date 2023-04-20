// https://www.hackerrank.com/challenges/separate-the-numbers/problem?isFullScreen=true
package main

import (
	"fmt"
)

func separateNumbers(s string) {
	for i := 1; i <= len(s)/2; i++ {
		first := s[:i]
		firstPrint := s[:i]

		j := i

		for {
			if j == len(s) {
				fmt.Println("YES", firstPrint)
				return
			}

			next := intToStr(1 + strToInt(first))
			if j+len(next) > len(s) || s[j:j+len(next)] != next {
				break
			}

			j += len(next)
			first = next
		}
	}

	fmt.Println("NO")
}

func strToInt(str string) int {
	numbers := int(0)

	// iterates string as rune
	for _, s := range str {
		// convert rune to integer by - '0'
		digit := int(s - '0')
		// multiply each digit with 10 to put
		// as the position and add the digit itself
		// e.g. 5 -> 5, 4 -> 40, 3 -> 300
		numbers = numbers*10 + digit
	}

	return numbers
}

func intToStr(n int) string {
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
	separateNumbers("123")
}
