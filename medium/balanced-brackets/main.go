// https://www.hackerrank.com/challenges/balanced-brackets/problem?isFullScreen=true
package main

import "fmt"

func isBalanced(s string) string {
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	// stack for storing the open bracket
	stack := make([]rune, 0)

	for _, bracket := range s {
		// if bracket is on the map, push to stack
		if _, ok := m[bracket]; ok {
			stack = append(stack, bracket)
		} else {
			// if stack is empty, it means the parentheses is invalid
			// e.g. ]} or []} then return false
			if len(stack) == 0 {
				return "NO"
			}

			// checks last index stack should be paired with current bracket
			// if not, return false
			// if paired, remove the last index
			lastIdx := len(stack) - 1
			if m[stack[lastIdx]] != bracket {
				return "NO"
			} else {
				stack = stack[:lastIdx]
			}
		}
	}

	if len(stack) == 0 {
		return "YES"
	}

	return "NO"
}

func main() {
	fmt.Println(isBalanced("{{[[(())]]}}"))
}
