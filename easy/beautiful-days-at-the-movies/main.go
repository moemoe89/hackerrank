package main

import "fmt"

func beautifulDays(i int32, j int32, k int32) int32 {
	var out int32

	for i = i; i <= j; i++ {
		reverseI := reverseInt32(int32ToString(i))
		subs := i - reverseI

		if subs%k == 0 {
			out++
		}
	}

	return out
}

func reverseInt32(n string) int32 {
	numbers := int32(0)

	for _, s := range n {
		digit := s - '0'

		numbers = numbers*10 + digit
	}

	reversed := int32(0)

	for numbers > 0 {
		reversed = reversed*10 + numbers%10

		numbers /= 10
	}

	return reversed
}

func int32ToString(n int32) string {
	var str []rune

	for n != 0 {
		m := n % 10

		r := '0' + rune(m)

		str = append([]rune{r}, str...)

		n /= 10
	}

	return string(str)
}

func main() {
	fmt.Println(beautifulDays(20, 23, 6))
}
