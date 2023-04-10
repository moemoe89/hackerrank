// https://www.hackerrank.com/challenges/strange-code/problem?isFullScreen=true
package main

import "fmt"

func strangeCounter(t int64) int64 {
	cycle := int64(3)

	for t > cycle {
		t -= cycle
		cycle *= 2
	}

	return cycle - t + 1
}

func main() {
	fmt.Println(strangeCounter(4))
}
