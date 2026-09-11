package partitionequalsubsetsum

func CanPartition(nums []int) bool {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {

		for s := target; s >= num; s-- {
			if dp[s-num] {
				dp[s] = true
			}
		}
		if dp[target] {
			return true
		}
	}

	return dp[target]
}
