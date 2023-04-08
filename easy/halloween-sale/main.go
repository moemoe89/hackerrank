// https://www.hackerrank.com/challenges/halloween-sale/problem?isFullScreen=true
package main

import "fmt"

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	count := int32(0)

	// Continue buying games while we have enough money
	for s >= p {
		count++         // Increment the number of games bought
		s -= p          // Deduct the price of the game from the remaining money
		p = max(p-d, m) // Update the price of the game (minimum price is m)
	}

	return count
}

func howManyGames2(p int32, d int32, m int32, s int32) int32 {
	if p > s {
		return 0
	}

	count := int32(1) // Buy the first game
	sum := p          // Keep track of the total price
	price := p        // Keep track of the current price

	for sum < s {
		// Compute the price of the next game
		price = max(price-d, m)

		if sum+price <= s {
			count++ // Buy the game if we have enough money
			sum += price
		} else {
			break // Stop buying games if we don't have enough money
		}
	}

	return count
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(howManyGames(100, 19, 1, 180))
}
