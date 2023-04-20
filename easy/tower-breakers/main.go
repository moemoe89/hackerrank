// https://www.hackerrank.com/challenges/tower-breakers-1/problem?isFullScreen=true
package main

import "fmt"

func towerBreakers(n int32, m int32) int32 {
	if m == 1 || n%2 == 0 {
		return 2
	}

	return 1
}

func main() {
	fmt.Println(towerBreakers(2, 6))
}
