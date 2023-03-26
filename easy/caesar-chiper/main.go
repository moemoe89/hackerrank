package main

import "fmt"

func caesarCipher(s string, k int32) string {
	// initialize number of alphabet
	n := 26

	// initialize array for alphabet
	arr := make([]rune, n)

	// put the alphabet into array
	for i := 0; i < n; i++ {
		arr[i] = rune('a' + i)
	}

	// the number of rotate for some test case can more than
	// the length of alphabet, it makes the rotation error.
	// for simplify the number, do modulo with the length of alphabet.
	k = k % int32(n)

	// rotate the array
	arr = append(arr[k:], arr[:k]...)

	// prepare encrypted array of rune.
	var encrypted []rune

	// iterates the string.
	for _, r := range s {
		// if the rune is lowercase alphabet
		if r >= 'a' && r <= 'z' {
			// convert to number of alphabet order
			// and minus 1 for match with the index.
			idx := int(r-'a'+1) - 1

			// replace rune value
			r = arr[idx]
		}

		// if the rune is uppercase alphabet.
		if r >= 'A' && r <= 'Z' {
			// convert to lowercase.
			r = r - 'A' + 'a'

			// convert to number of alphabet order
			// and minus 1 for match with the index.
			idx := int(r-'a'+1) - 1

			// replace rune value.
			r = arr[idx]

			// convert back to uppercase.
			r = r - 'a' + 'A'
		}

		// put to encrypted array of rune.
		encrypted = append(encrypted, r)
	}

	return string(encrypted)
}

func main() {
	fmt.Println(caesarCipher("www.abc.xy", 87))
}
