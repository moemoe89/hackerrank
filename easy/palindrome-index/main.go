// https://www.hackerrank.com/challenges/palindrome-index/problem?isFullScreen=true
package main

import "fmt"

func palindromeIndex(s string) int32 {
	// initialize pointers to the beginning and end of the string
	i := 0
	j := len(s) - 1

	// iterate over the string from both ends
	for i < j {
		// if the characters at the current positions do not match, check if
		// removing either of them will result in a palindrome
		if s[i] != s[j] {
			// check if removing the character at index i will result in a palindrome
			if isPalindrome(s[:i] + s[i+1:]) {
				return int32(i)
			}
			// check if removing the character at index j will result in a palindrome
			if isPalindrome(s[:j] + s[j+1:]) {
				return int32(j)
			}
			// if neither option results in a palindrome, return -1
			return -1
		}
		// move pointers towards the center of the string
		i++
		j--
	}
	// if we reach the center of the string and haven't returned yet, it's already a palindrome
	return -1
}

func isPalindrome(s string) bool {
	// initialize pointers to the beginning and end of the string
	i := 0
	j := len(s) - 1

	// iterate over the string from both ends
	for i < j {
		// if the characters at the current positions do not match, the string is not a palindrome
		if s[i] != s[j] {
			return false
		}
		// move pointers towards the center of the string
		i++
		j--
	}
	// if we reach the center of the string and haven't returned yet, the string is a palindrome
	return true
}

func main() {
	fmt.Println(palindromeIndex("aaab"))
}
