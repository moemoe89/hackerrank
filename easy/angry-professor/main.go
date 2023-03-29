package main

import "fmt"

func angryProfessor(k int32, a []int32) string {
	// initialize the numbers of students.
	n := int32(len(a))

	// variable for counting the late students.
	var numLate int32

	for _, v := range a {
		// if the value more than 0,
		// it's mean the student is late.
		if v > 0 {
			numLate++
		}
	}

	// count the student that available to attend.
	attend := n - numLate

	// if the number of attend more than the expected
	// the class will go on.
	if attend < k {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Println(angryProfessor(3, []int32{-2, -1, 0, 1, 2}))
}
