// https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true
package main

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	output := int32(0)

	for i := a[len(a)-1]; i <= b[0]; i++ {
		isFactorOfA := true

		for _, val := range a {
			if i%val != 0 {
				isFactorOfA = false
				break
			}
		}

		if !isFactorOfA {
			continue
		}

		// Check if all elements in b are factors of i
		isFactorOfB := true

		for _, val := range b {
			if val%i != 0 {
				isFactorOfB = false
				break
			}
		}

		if isFactorOfB {
			output++
		}
	}

	return output
}

func main() {
	fmt.Println(getTotalX([]int32{2, 6}, []int32{24, 36}))
}
