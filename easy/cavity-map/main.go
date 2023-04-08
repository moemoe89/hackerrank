// https://www.hackerrank.com/challenges/cavity-map/problem?isFullScreen=true
package main

import "fmt"

func cavityMap(grid []string) []string {
	n := len(grid[0]) - 1

	result := grid

	for i := range grid {
		if i == 0 || i == n {
			continue
		}

		runes := []rune(grid[i])

		for j := 1; j < n; j++ {
			cur := grid[i][j]
			left := grid[i][j-1]
			right := grid[i][j+1]
			upper := grid[i-1][j]
			lower := grid[i+1][j]

			if cur > left && cur > right && cur > upper && cur > lower {
				runes[j] = 'X'
			}
		}

		result[i] = string(runes)
	}

	return result
}

func main() {
	fmt.Println(cavityMap([]string{"1112", "1912", "1892", "1234"}))
}
