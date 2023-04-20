// https://www.hackerrank.com/challenges/acm-icpc-team/problem?isFullScreen=true
package main

import "fmt"

func acmTeam(topic []string) []int32 {
	maxTopics := int32(0)

	m := make(map[int32]int32)

	for i := 0; i < len(topic)-1; i++ {

		for j := i + 1; j < len(topic); j++ {
			topics := int32(0)

			for k := 0; k < len(topic[i]); k++ {

				if topic[i][k] == '1' || topic[j][k] == '1' {
					topics++
				}

			}

			m[topics]++

			if topics > maxTopics {
				maxTopics = topics
			}
		}

	}

	return []int32{maxTopics, m[maxTopics]}
}

func main() {
	fmt.Println(acmTeam([]string{
		"10101",
		"11100",
		"11010",
		"00101",
	}))
}
