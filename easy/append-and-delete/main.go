// https://www.hackerrank.com/challenges/append-and-delete/problem?isFullScreen=true
package main

import "fmt"

func appendAndDelete(s string, t string, k int32) string {
	lenS, lenT := len(s), len(t)
	if lenS+lenT <= int(k) {
		return "Yes"
	}

	samePrefix := 0
	for samePrefix < lenS && samePrefix < lenT && s[samePrefix] == t[samePrefix] {
		samePrefix++
	}

	deleteS := lenS - samePrefix
	deleteT := lenT - samePrefix

	remainingOps := k - int32(deleteS+deleteT)
	if remainingOps >= 0 && remainingOps%2 == 0 {
		return "Yes"
	}

	return "No"
}

func main() {
	fmt.Println(appendAndDelete("abcd", "abcdert", 6))
}
