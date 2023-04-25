// https://www.hackerrank.com/challenges/strange-advertising/problem?isFullScreen=true
package main

import "fmt"

func viralAdvertising(n int32) int32 {
	var out, likes, i int32

	out = 2
	likes = 2

	for i = 2; i <= n; i++ {
		adv := likes * 3
		halfAdv := adv / 2
		out += halfAdv
		likes = halfAdv
	}

	return out
}

func main() {
	fmt.Println(viralAdvertising(3))
}
