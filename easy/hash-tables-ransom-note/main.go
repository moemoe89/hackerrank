// https://www.hackerrank.com/challenges/ctci-ransom-note/problem?isFullScreen=true
package main

import "fmt"

func checkMagazine(magazine []string, note []string) {
	m := make(map[string]int, len(magazine))

	for _, v := range magazine {
		m[v]++
	}

	for _, n := range note {
		if _, ok := m[n]; !ok {
			fmt.Println("No")
			return
		}

		if m[n] == 0 {
			fmt.Println("No")
			return
		}

		m[n]--
	}

	fmt.Println("Yes")
}

func main() {
	checkMagazine([]string{"two", "times", "three", "is", "not", "four"}, []string{"two", "times", "two", "is", "four"})
}
