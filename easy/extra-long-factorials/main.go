// https://www.hackerrank.com/challenges/extra-long-factorials/problem?isFullScreen=true
package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int64) {
	out := big.NewInt(n)
	for n = n - 1; n >= 1; n-- {
		out = new(big.Int).Mul(out, big.NewInt(n))
	}

	fmt.Println(out)
}

func main() {
	extraLongFactorials(30)
}
