package main

import "fmt"

func catAndMouse(x int32, y int32, z int32) string {
	distanceY := z - y
	if z < y {
		distanceY = y - z
	}

	distanceX := z - x
	if z < x {
		distanceX = x - z
	}

	if distanceY == distanceX {
		return "Mouse C"
	}

	if distanceX < distanceY {
		return "Cat A"
	}

	return "Cat B"
}

func main() {
	fmt.Println(catAndMouse(1, 2, 3))
}
