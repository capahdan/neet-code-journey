package besttimetobuystockwithcooldown

func MaxProfit(prices []int) int {
	// state: buying or selling?
	// if buy  -> i + 1
	// if sell -> i + 2
	type state struct {
		i      int
		buying bool
	}

	dp := make(map[state]int) // key (i, buying) -> max profit

	var dfs func(i int, buying bool) int
	dfs = func(i int, buying bool) int {
		// base case: past the end (also covers empty slice)
		if i >= len(prices) {
			return 0
		}
		key := state{i, buying}
		if v, ok := dp[key]; ok {
			return v
		}

		cooldown := dfs(i+1, buying)
		var best int
		if buying {
			buy := dfs(i+1, !buying) - prices[i]
			best = max(buy, cooldown)
		} else {
			sell := dfs(i+2, !buying) + prices[i]
			best = max(sell, cooldown)
		}

		dp[key] = best
		return best
	}

	return dfs(0, true)
}
