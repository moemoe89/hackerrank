// https://www.hackerrank.com/challenges/maximum-element/problem?isFullScreen=true
package main

import "fmt"

func getMax(operations []string) []int32 {
	var output []int32

	var stack, maxStack []int32

	for _, o := range operations {
		query := o[0]

		if query == '1' {
			n := strToInt32(o[2:])

			stack = append(stack, n)
			if len(maxStack) == 0 || n >= maxStack[len(maxStack)-1] {
				maxStack = append(maxStack, n)
			}
		} else if query == '2' {
			if len(stack) == 0 {
				continue
			}

			if stack[len(stack)-1] == maxStack[len(maxStack)-1] {
				maxStack = maxStack[:len(maxStack)-1]
			}

			stack = stack[:len(stack)-1]
		} else if query == '3' {
			if len(stack) == 0 {
				continue
			}

			output = append(output, maxStack[len(maxStack)-1])
		}
	}

	return output
}

func strToInt32(str string) int32 {
	numbers := int32(0)

	// iterates string as rune
	for _, s := range str {
		// convert rune to integer by - '0'
		digit := int(s - '0')
		// multiply each digit with 10 to put
		// as the position and add the digit itself
		// e.g. 5 -> 5, 4 -> 40, 3 -> 300
		numbers = numbers*10 + int32(digit)
	}

	return numbers
}

func main() {
	fmt.Println(getMax([]string{"1 97", "2", "1 20", "2", "1 26", "1 20", "2", "3", "1 19", "3"}))
}
