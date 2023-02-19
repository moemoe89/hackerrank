package main

import "fmt"

func marsExploration(s string) int32 {
	var out int32

	var j int

	for i, v := range s {
		if i%3 == 0 {
			j = 0
		}

		if j == 0 && v != 'S' {
			out++
		} else if j == 1 && v != 'O' {
			out++
		} else if j == 2 && v != 'S' {
			out++
		}

		j++
	}

	return out
}

func main() {
	fmt.Println(marsExploration("SOSSPSSQSSOR"))
}
