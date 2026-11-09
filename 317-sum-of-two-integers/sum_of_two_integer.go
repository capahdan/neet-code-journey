package sumoftwointegers

func GetSum(a, b int) int {
	x, y := int32(a), int32(b)
	var carry int32

	for y != 0 {
		carry = x & y
		x = x ^ y
		y = carry << 1
	}
	return int(x)
}
