package longestpalindromicsubstring

func longestPalindrome(s string) string {
	result := ""

	var count func(l, r int)
	count = func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if len(result) < len(s[l:r+1]) {
				result = s[l : r+1]
			}
			l--
			r++
		}
	}

	for i := range s {
		count(i, i)
		count(i, i+1)
	}

	return result
}
