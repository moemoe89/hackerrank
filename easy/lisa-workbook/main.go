// https://www.hackerrank.com/challenges/lisa-workbook/problem?isFullScreen=true
package main

import "fmt"

func workbook(n int32, k int32, arr []int32) int32 {
	pages := int32(0)

	special := int32(0)

	for i := range arr {
		addPage := int32(0)

		if arr[i] > k {
			addPage += arr[i] / k

			if arr[i]%k != 0 {
				addPage++
			}
		} else {
			addPage++
		}

		numProblem := int32(0)
		curProblem := arr[i]

		l := int32(1)

		for j := pages + 1; j <= pages+addPage; j++ {
			minProblem := numProblem + 1

			maxProblem := k * l

			if curProblem < k {
				maxProblem = numProblem + curProblem
			}

			if j >= minProblem && j <= maxProblem {
				special++
			}

			numProblem = maxProblem
			curProblem -= k
			l++
		}

		pages += addPage
	}

	return special
}

func main() {
	fmt.Println(workbook(5, 3, []int32{4, 2, 6, 1, 10}))
}
