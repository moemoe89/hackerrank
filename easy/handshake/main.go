// https://www.hackerrank.com/challenges/handshake/problem?isFullScreen=true
package handshake

import "fmt"

func handshake(n int32) int32 {
	if n < 2 {
		return 0
	}

	return n*(n+1)/int32(2) - n
}

func main() {
	fmt.Println(handshake(10))
}
