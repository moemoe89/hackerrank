package main

import "fmt"

func timeConversion(s string) string {
	n := len(s)

	r := []rune(s)

	if s[n-2] == 'P' {
		x := int(s[0]-'0') * 10
		x2 := (int(s[1] - '0')) + x

		if x2 == 12 {
			return string(r[:n-2])
		}

		h := 0

		if int(s[0]-'0') > 0 {

			h = int(s[0]-'0') * 10
			h += int(s[1] - '0')
			h += 12

		} else {
			h = int(s[1] - '0')
			h += 12
		}

		r[0] = '0' + rune(h/10)
		r[1] = '0' + rune(h%10)

	} else {
		x := int(s[0]-'0') * 10
		x2 := (int(s[1] - '0')) + x

		if x2 == 12 {
			r[0] = '0' + rune(0)
			r[1] = '0' + rune(0)
		}
	}

	return string(r[:n-2])
}

func main() {
	fmt.Println(timeConversion("07:05:45PM"))
}
