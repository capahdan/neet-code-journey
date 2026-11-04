package regulerexpressionmatching

type item struct {
	i int
	j int
}

func IsMatch(s, p string) bool {

	dp := make(map[item]bool)
	var dfs func(i, j int) bool

	dfs = func(i, j int) bool {
		key := item{
			i: i,
			j: j,
		}

		if val, ok := dp[key]; ok {
			return val
		}
		if i >= len(s) && j >= len(p) {
			return true
		}

		if j >= len(p) {
			return false
		}

		match := i < len(s) && (s[i] == p[j] || p[j] == '.')

		if j+1 < len(p) && p[j+1] == '*' {
			// we can choose whether to use or not using prev character
			dp[key] = dfs(i, j+2) || (match && dfs(i+1, j))
			return dp[key]
		}

		if match {
			dp[key] = dfs(i+1, j+1)
			return dp[key]
		}

		dp[key] = false
		return dp[key]
	}

	return dfs(0, 0)
}
