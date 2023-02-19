package main

import "fmt"

func designerPdfViewer(h []int32, word string) int32 {
	var out int32

	mapChar := make(map[rune]int32)

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i, r := range alphabet {
		mapChar[r] = h[i]
	}

	for i, r := range word {
		if i == 0 {
			out = mapChar[r]
		} else {
			if mapChar[r] > out {
				out = mapChar[r]
			}
		}
	}

	return out * int32(len(word))
}

func main() {
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, "abc"))
}
