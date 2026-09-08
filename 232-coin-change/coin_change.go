package coinchange

import "math"

// the intution behind this algorithm is that
// we don't need to count all the combination of coins in one go
// we only need remember what is the calculation before
// like for dp[i] we just need to check the dp[i-coin] and add 1
// after that we just need to calculate the min(dp[i-coin]+1) and take it
// we start by adding amount from 1 to amount
// after that we just check all the coins we have
// if we have coins that smaller than i we check whether
// the dp[i-coin] + 1 is smaller then the dp[i]
// if so we update the current dp

// func coinchange(coins []int, amount int) int {
// 	if amount == 0 {
// 		return 0
// 	}

// 	dp := make([]int, amount+1)
// 	for i := 1; i <= amount; i++ {
// 		dp[i] = amount + 1
// 	}

// 	dp[0] = 0

// 	for i := 1; i <= amount; i++ {
// 		for _, coin := range coins {
// 			if coin <= i {
// 				if dp[i-coin]+1 < dp[i] {
// 					dp[i] = dp[i-coin] + 1
// 				}
// 			}
// 		}
// 	}

// 	if dp[amount] > amount {
// 		return -1
// 	}
// 	return dp[amount]
// }

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	var minCount = math.MaxInt

	for _, coin := range coins {
		res := coinChange(coins, amount-coin)

		if res >= 0 && res < minCount {
			minCount = res + 1
		}
	}

	if minCount == math.MaxInt {
		return -1
	}
	return minCount
}
