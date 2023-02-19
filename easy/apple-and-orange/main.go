package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	var firstOutput, secondOutput int
	for _, apple := range apples {
		if (a+apple) >= s && (a+apple) <= t {
			firstOutput++
		}
	}

	for _, orange := range oranges {
		if (b+orange) >= s && (b+orange) <= t {
			secondOutput++
		}
	}

	fmt.Println(firstOutput)
	fmt.Println(secondOutput)
}

func main() {
	countApplesAndOranges(5, 15, 3, 2, []int32{-2, 2, 1}, []int32{5, -6})
}
