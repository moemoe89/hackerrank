package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// prepare the number of apple and orange falls.
	var appleFalls, orangeFalls int

	// check if each apple in the Sam's house location or not (s <= x <= t).
	// need to add apple tree location (a)
	for _, apple := range apples {
		if (a+apple) >= s && (a+apple) <= t {
			appleFalls++
		}
	}

	// check if each orange in the Sam's house location or not (s <= x <= t).
	// need to add orange tree location (b)
	for _, orange := range oranges {
		if (b+orange) >= s && (b+orange) <= t {
			orangeFalls++
		}
	}

	fmt.Println(appleFalls)
	fmt.Println(orangeFalls)
}

func main() {
	countApplesAndOranges(5, 15, 3, 2, []int32{-2, 2, 1}, []int32{5, -6})
}
