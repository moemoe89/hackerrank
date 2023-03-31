// https://www.hackerrank.com/challenges/bon-appetit/problem?isFullScreen=true
package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	var total int32

	// calculate the total cost.
	for _, v := range bill {
		total += v
	}

	// calculate the split in half
	splitBill := (total - bill[k]) / 2
	// calculate the rest between charged and split bill.
	rest := b - splitBill

	// if zero, print Bon Appetit.
	if rest == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(rest)
	}
}

func main() {
	bonAppetit([]int32{2, 4, 6}, 2, 10)
}
