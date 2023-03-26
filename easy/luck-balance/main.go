package main

import "fmt"

func luckBalance(k int32, contests [][]int32) int32 {
	// descending sort the contests by first index
	bubbleSort2D(contests)

	var output int32

	// iterates contests
	for i := range contests {
		// if the contest not important,
		// always accumulated the output
		if contests[i][1] == 0 {
			output += contests[i][0]

			continue
		}

		// if the contest is important, will go from here.
		//
		// if we still have the luck, we can use it
		// then decrease the luck because being used.
		if k > 0 {
			output += contests[i][0]
			k--

			continue
		}

		// if there's no luck, this contest will lose
		// and minus the value.
		output -= contests[i][0]
	}

	return output
}

func bubbleSort2D(arr [][]int32) {
	n := len(arr) - 1
	for i := 0; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			if arr[j][0] < arr[j+1][0] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	fmt.Println(luckBalance(2, [][]int32{{5, 1}, {1, 1}, {4, 0}}))
}
