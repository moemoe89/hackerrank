// https://www.hackerrank.com/challenges/day-of-the-programmer/problem?isFullScreen=true
package main

import "fmt"

func dayOfProgrammer(year int32) string {
	if year == 1918 {
		// In 1918, the 256th day of the year was September 26th, as there was a calendar change
		return "26.09.1918"
	} else if (year <= 1917 && year%4 == 0) || (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
		// Julian or Gregorian leap year, so the 256th day is on September 12th
		return "12.09." + int32toStr(year)
	} else {
		// Not a leap year, so the 256th day is on September 13th
		return "13.09." + int32toStr(year)
	}
}

func int32toStr(numbers int32) string {
	var str []rune

	// iterates until numbers become 0
	for numbers != 0 {
		// use modulo 10 in order to get every digit
		// e.g. 12345 % 10 -> 5
		m := numbers % 10

		// convert to rune string
		r := '0' + rune(m)

		// use prepend because we need to add
		// digit from right to left
		// [5] -> [4 5] -> [3 4 5]
		str = append([]rune{r}, str...)

		// divide the number to eliminate the last digit
		// 12345 to 1234
		numbers /= 10
	}

	return string(str)
}

func main() {
	fmt.Println(dayOfProgrammer(1989))
}
