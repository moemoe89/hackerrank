package main

import "fmt"

func repeatedString(s string, n int64) int64 {
	output := int64(0)

	missing := n % int64(len(s))

	missingOutput := int64(0)

	for i := range s {
		if 'a' == s[i] {
			output++
		}

		if int64(i) < missing && 'a' == s[i] {
			missingOutput++
		}
	}

	numRepeats := n / int64(len(s))

	aCountRepeated := output * numRepeats

	return aCountRepeated + missingOutput
}

func main() {
	fmt.Println(repeatedString("a", 1))
}
