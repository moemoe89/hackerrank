// https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=true
package main

import "fmt"

func angryProfessor(k int32, a []int32) string {
	// variable for counting the not late students.
	var attend int32

	for _, v := range a {
		// if the value under equal than 0,
		// it's mean the student is not late.
		if v <= 0 {
			attend++
		}
	}

	// if the number of attend more than the expected
	// the class will not cancelled.
	if attend >= k {
		return "NO"
	}

	return "YES"
}

func main() {
	fmt.Println(angryProfessor(3, []int32{-2, -1, 0, 1, 2}))
}
