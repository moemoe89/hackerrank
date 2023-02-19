package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if (x1 > x2 && v1 > v2) || (x1 < x2 && v1 < v2) || ((v2 - v1) == 0) {
		return "NO"
	}

	if x1 < x2 && v1 < v2 {
		return "NO"
	}

	if (x1-x2)%(v2-v1) == 0 {
		return "YES"
	}

	return "NO"
}

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2))
}
