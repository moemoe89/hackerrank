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

	numOps := int32(deleteS + deleteT)
	if numOps <= k && (k-numOps)%2 == 0 {
		return "Yes"
	}

	return "No"
}

func main() {
	fmt.Println(appendAndDelete("abc", "def", 6))
}
