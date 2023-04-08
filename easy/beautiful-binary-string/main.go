// https://www.hackerrank.com/challenges/beautiful-binary-string/problem?isFullScreen=true
package main

import "fmt"

func beautifulBinaryString(b string) int32 {
	output := int32(0)

	for i := 0; i <= len(b)-3; i++ {
		if string(b[i:i+3]) == "010" {
			output++

			i += 2
		}
	}

	return output
}

func main() {
	fmt.Println(beautifulBinaryString("0101010"))
}
