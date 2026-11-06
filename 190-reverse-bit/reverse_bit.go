package reversebit

func reverseBit(n int) int {
	arr := parseToBinaryArray(n)

	return calculate(arr)
}

func parseToBinaryArray(n int) []int {
	arr := make([]int, 32)
	for i := 0; i < 32; i++ {
		arr[i] = n % 2
		n /= 2
	}
	return arr
}

func calculate(arr []int) int {
	res := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] == 1 {
			res += 1 << (n - 1 - i)
		}
	}

	return res
}
