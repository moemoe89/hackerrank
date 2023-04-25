// https://www.hackerrank.com/challenges/library-fine/problem?isFullScreen=true
package main

import "fmt"

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1 > y2 {
		return 10000
	} else if y1 < y2 {
		return 0
	}

	if m1 > m2 {
		return (m1 - m2) * 500
	} else if m1 < m2 {
		return 0
	}

	if d1 > d2 {
		return (d1 - d2) * 15
	}

	return 0
}

func main() {
	fmt.Println(libraryFine(9, 6, 2015, 6, 6, 2015))
}
