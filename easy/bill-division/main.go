package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	var total int32

	for _, v := range bill {
		total += v
	}

	splitBill := (total - bill[k]) / 2
	rest := b - splitBill

	if rest == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(rest)
	}
}

func main() {
	bonAppetit([]int32{2, 4, 6}, 2, 10)
}
