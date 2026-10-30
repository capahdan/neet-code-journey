package coinchangeii

// func Change(amount int, coins []int) int {

// 	dp := make([]int, amount+1)
// 	dp[0] = 1

// 	for _, coin := range coins {
// 		for a := coin; a <= amount; a++ {
// 			dp[a] += dp[a-coin]
// 		}
// 	}

// 	return dp[amount]
// }

type key struct {
	i int
	a int
}

func Change(amount int, coins []int) int {
	cache := make(map[key]int)

	var dfs func(i, a int) int
	dfs = func(i, a int) int {
		if a == amount {
			return 1
		}

		if a > amount {
			return 0
		}

		if i == len(coins) {
			return 0
		}

		k := key{
			i: i,
			a: a,
		}
		if v, ok := cache[k]; ok {
			return v
		}

		cache[k] = dfs(i, a+coins[i]) + dfs(i+1, a)
		return cache[k]
	}

	return dfs(0, 0)
}
