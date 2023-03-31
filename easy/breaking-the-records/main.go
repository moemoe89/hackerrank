// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem?isFullScreen=true
package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	// the first game, min & max count not counted,
	// then set it by 0.
	minCount, maxCount := int32(0), int32(0)

	// the first game always be min and max score,
	// then initialize the value based on first game score.
	minScore, maxScore := scores[0], scores[0]

	// iterates the scores.
	for i := 1; i < len(scores); i++ {
		// if the current score more than max score
		// increment the max count and replace the max score.
		if scores[i] > maxScore {
			maxScore = scores[i]
			maxCount++
		}

		// if the current score lower than min score
		// increment the min count and replace the min score.
		if scores[i] < minScore {
			minScore = scores[i]
			minCount++
		}
	}

	return []int32{maxCount, minCount}
}

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
}
