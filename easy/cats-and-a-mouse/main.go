// https://www.hackerrank.com/challenges/cats-and-a-mouse/problem?isFullScreen=true
package main

import "fmt"

func catAndMouse(x int32, y int32, z int32) string {
	// check distance x (Cat A), return as abs number.
	// for the example target in 2, and x position in 3.
	// 2-3 = -1, then return as 1.
	distanceX := abs(z - x)
	// check distance y (Cat B), return as abs number.
	distanceY := abs(z - y)

	// if the distance same, return Mouse C
	if distanceX == distanceY {
		return "Mouse C"
	}

	// return cat A if nearest to the target.
	if distanceX < distanceY {
		return "Cat A"
	}

	return "Cat B"
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}

	return n
}

func main() {
	fmt.Println(catAndMouse(1, 2, 3))
}
