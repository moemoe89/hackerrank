// https://www.hackerrank.com/challenges/staircase/problem?isFullScreen=true
package main

import "fmt"

func staircase(n int32) {
	for i := int32(0); i < n; i++ {

		for j := int32(0); j < n; j++ {
			if j >= n-i-1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println("")
	}
}

func main() {
	staircase(6)
}
