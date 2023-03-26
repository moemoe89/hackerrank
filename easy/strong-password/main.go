package main

import (
	"fmt"
)

func minimumNumber(n int32, password string) int32 {
	// initialize the min number to make the password strong.
	minNum := int32(0)

	// prepare the flag for make the password strong.
	containsDigit := false
	containsLowercase := false
	containsUppercase := false
	containsSpecialChar := false

	// iterates the password
	for _, r := range password {
		// check if the password contains digit.
		if r >= '0' && r <= '9' {
			containsDigit = true
		}

		// check if the password contains uppercase.
		if r >= 'A' && r <= 'Z' {
			containsUppercase = true
		}

		// check if the password contains lowercase.
		if r >= 'a' && r <= 'z' {
			containsLowercase = true
		}

		// check if the password contains special characters.
		if isPunctuation(r) {
			containsSpecialChar = true
		}
	}

	// if not contains the required factor to make
	// password strong, increment the min number

	if !containsDigit {
		minNum++
	}

	if !containsLowercase {
		minNum++
	}

	if !containsUppercase {
		minNum++
	}

	if !containsSpecialChar {
		minNum++
	}

	// if password less than 6,
	// if the min number to add already greater that the missing length,
	// we can just use the min number, but if it doesn't, we will use the missing length number.
	// for the example password is Ab1, need more 3 characters also from the factor we need only special char.
	// but if we only add special character, we are still missing the length, so it should use the missing length.
	if len(password) < 6 {
		minNum = max(minNum, 6-int32(len(password)))
	}

	return minNum
}

func isPunctuation(r rune) bool {
	punctuation := `!@#$%^&*()-+`

	for _, p := range punctuation {
		if p == r {
			return true
		}
	}

	return false
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(minimumNumber(3, "Ab1"))
}
