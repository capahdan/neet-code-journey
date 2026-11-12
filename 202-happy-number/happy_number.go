package happynumber

func IsHappy(n int) bool {
	seen := make(map[int]int)

	run := true
	for run {
		if n == 1 {
			return true
		}

		val, ok := seen[n]
		if ok && val > 1 {
			break
		}
		arr := parseToArray(n)
		n = parseArrayToInt(arr)
		seen[n] += 1
	}

	return false
}

func parseArrayToInt(arr []int) int {
	res := 0
	for _, num := range arr {
		res += num * num
	}
	return res
}

func parseToArray(n int) []int {
	result := []int{}

	for n != 0 {
		modulo := n % 10
		result = append(result, modulo)
		n /= 10
	}

	return result
}
