package main

import "fmt"

func luckBalance(k int32, contests [][]int32) int32 {
	bubbleSort2D(contests)

	var output int32

	for i := range contests {
		if contests[i][1] == 0 {
			output += contests[i][0]

			continue
		}

		if k > 0 {
			output += contests[i][0]
			k--

			continue
		}

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
