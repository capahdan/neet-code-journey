package plusone

func PlusOne(arr []int) []int {

	lastDigit := len(arr) - 1
	if arr[lastDigit] != 9 {
		arr[lastDigit] += 1
		return arr
	}
	arr[lastDigit] += 1
	carry := 0

	for i := lastDigit; i > 0; i-- {
		carry = arr[i] / 10
		arr[i] = arr[i] % 10
		arr[i-1] += carry
	}

	if arr[0] == 10 {
		newArr := make([]int, len(arr)+1)
		newArr[0] = 1
		return newArr
	}

	return arr

}
