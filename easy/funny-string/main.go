// https://www.hackerrank.com/challenges/funny-string/problem?isFullScreen=true
package main

import "fmt"

func funnyString(s string) string {
	n := len(s) - 1

	for i := range s {
		if i < n {
			diff1 := int(s[i+1]) - int(s[i])

			diff2 := int(s[n-i]) - int(s[n-i-1])

			if diff1 < 0 {
				diff1 *= -1
			}

			if diff2 < 0 {
				diff2 *= -1
			}

			if diff1 != diff2 {
				return "Not Funny"
			}
		}
	}

	return "Funny"
}

func main() {
	fmt.Println(funnyString("acxz"))
}
