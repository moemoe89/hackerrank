// https://www.hackerrank.com/challenges/find-digits/problem?isFullScreen=true
package main

import "fmt"

func findDigits(n int32) int32 {
	output := int32(0)

	num := n

	for n != 0 {
		m := n % 10

		if m != 0 && num%m == 0 {
			output++
		}

		n /= 10
	}

	return output
}

func main() {
	fmt.Println(findDigits(1012))
}
