package main

import "fmt"

func checkMagazine(magazine []string, note []string) {
	m := make(map[string]int, len(magazine))

	for i := range magazine {
		m[magazine[i]]++
	}

	for i := range note {
		if _, ok := m[note[i]]; !ok {
			fmt.Println("No")
			return
		}

		if m[note[i]] == 0 {
			fmt.Println("No")
			return
		}

		m[note[i]]--
	}

	fmt.Println("Yes")
}

func main() {
	checkMagazine([]string{"two", "times", "three", "is", "not", "four"}, []string{"two", "times", "two", "is", "four"})
}
