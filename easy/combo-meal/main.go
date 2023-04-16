// https://www.hackerrank.com/challenges/combo-meal/problem?isFullScreen=true
package main

import "fmt"

func profit(b int32, s int32, c int32) int32 {
	return b + s - c
}

func main() {
	fmt.Println(profit(6, 9, 11))
}
