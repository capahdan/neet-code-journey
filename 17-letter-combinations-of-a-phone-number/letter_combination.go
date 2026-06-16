package lettercombinationsofaphonenumber

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	letterMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var result []string
	var backTrack func(index int, path string)
	backTrack = func(index int, path string) {
		if len(path) == len(digits) {
			result = append(result, path)
			return
		}

		for _, letter := range letterMap[string(digits[index])] {
			backTrack(index+1, path+string(letter))
		}
	}

	backTrack(0, "")
	return result
}
