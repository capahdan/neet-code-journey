package hummingweight

func hammingWeight(n int) int {
	count := 0

	for n > 0 {
		mod := n % 2
		if mod == 1 {
			count += 1
		}
		n = n / 2
	}
	return count
}

// bit mask

func hammingWeightBitmax(n int) int {
	res := 0
	for n != 0 {
		if n&1 != 0 {
			res++
		}
	}
	n >>= 1
	return res
}

func hammingWeightOptimal(n int) int {
	res := 0
	for n != 0 {
		n &= n - 1
		res++
	}
	return res

}
