package palindromicsubstring

// the intuition behind this function is that we will go each character from the left to the right
// then we and one caracther to the left one and to the right one if it palindrom we count the result
// and also we should consider even substring
func CountSubstring(s string) int {
	res := 0

	var count func(l, r int)
	count = func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}

	for i := range s {
		count(i, i)
		count(i, i+1)
	}

	return res
}
