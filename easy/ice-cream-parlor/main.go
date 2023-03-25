package main

import "fmt"

func icecreamParlor(m int32, arr []int32) []int32 {
	mArr := make(map[int32]int, len(arr))

	for i, a := range arr {
		if val, ok := mArr[m-a]; ok {
			return []int32{int32(val + 1), int32(i + 1)}
		}

		mArr[a] = i
	}

	return []int32{}
}

func icecreamParlor2(m int32, arr []int32) []int32 {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[i]+arr[j] == m {
				return []int32{int32(i + 1), int32(j + 1)}
			}
		}
	}

	return []int32{}
}

func main() {
	fmt.Println(icecreamParlor(6, []int32{1, 3, 4, 5, 6}))
}
