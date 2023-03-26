package main

import (
	"fmt"
)

// use int64 instead because some test case have a big numbers.
func kaprekarNumbers(p int64, q int64) {
	// initialize the output array.
	var output []int64

	// iterates from p to q.
	for i := p; i <= q; i++ {
		// 1 should be counted, then add checker here.
		if i == 1 {
			output = append(output, i)

			continue
		}

		// do square calculation (i^2)
		square := i * i

		// if the digit lower than 10, skip it.
		if square < 10 {
			continue
		}

		// convert to string.
		squareStr := numbersToStr(square)

		// count the length of string.
		n := len(squareStr)

		// split into two numbers.
		left := squareStr[:n/2]
		right := squareStr[n/2:]

		// convert to int64
		leftNumber := strToNumbers(left)
		rightNumber := strToNumbers(right)

		// if the number i equal with left + right,
		// it''s kaprekar number.
		if i == (leftNumber + rightNumber) {
			output = append(output, i)
		}
	}

	// if no kaprekar number found, print invalid.
	if len(output) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		for _, o := range output {
			fmt.Printf("%d ", o)
		}
	}
}

func numbersToStr(n int64) string {
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

func strToNumbers(str string) int64 {
	numbers := int64(0)

	for _, s := range str {
		// convert rune to integer by - '0'
		digit := int(s - '0')
		// multiply each digit with 10 to put
		// as the position and add the digit itself
		// e.g. 5 -> 5, 4 -> 40, 3 -> 300
		numbers = numbers*int64(10) + int64(digit)
	}

	return numbers
}

func main() {
	kaprekarNumbers(1, 99999)
}
