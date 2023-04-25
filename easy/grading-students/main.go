// https://www.hackerrank.com/challenges/grading/problem?isFullScreen=true
package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	for i := range grades {
		if grades[i] < 38 {
			continue
		}

		target := grades[i] / 5
		target *= 5
		target += 5

		if target-grades[i] < 3 {
			grades[i] = target
		}
	}

	return grades
}

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))
}
