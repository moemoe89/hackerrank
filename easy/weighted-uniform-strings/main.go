// https://www.hackerrank.com/challenges/weighted-uniform-string/problem?isFullScreen=true
package main

import "fmt"

func weightedUniformStrings(s string, queries []int32) []string {
	weights := make(map[int32]struct{})

	var lastChar rune
	var weight int32

	for _, r := range s {
		if r == lastChar {
			weight += r - 'a' + 1
		} else {
			weight = r - 'a' + 1
		}

		lastChar = r

		weights[weight] = struct{}{}
	}

	outputs := make([]string, len(queries))

	for i, q := range queries {
		output := "Yes"

		if _, ok := weights[q]; !ok {
			output = "No"
		}

		outputs[i] = output
	}

	return outputs
}

func main() {
	fmt.Println(weightedUniformStrings("abccddde", []int32{1, 3, 12, 5, 9, 10}))
}
