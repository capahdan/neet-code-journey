package palindromepartitioning

func partition(s string) [][]string {
	var result [][]string
	var current []string

	var backTrack func(start int)
	backTrack = func(start int) {
		if start == len(s) {
			temp := make([]string, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if isPalindrom(sub) {
				current = append(current, sub)
				backTrack(end)
				current = current[:len(current)-1]
			}
		}

	}

	backTrack(0)

	return result
}

func isPalindrom(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
