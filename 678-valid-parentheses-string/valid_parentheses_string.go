package validparenthesesstring

func IsValidString(s string) bool {

	leftStack := []rune{}
	starStack := []rune{}

	for _, c := range s {
		switch c {
		case '(':
			leftStack = append(leftStack, c)
		case '*':
			starStack = append(starStack, c)
		case ')':

			if len(leftStack) > 0 {
				leftStack = leftStack[:len(leftStack)-1]
			} else if len(starStack) > 0 {
				starStack = starStack[:len(starStack)-1]
			} else {
				return false
			}
		}
	}

	for len(leftStack) > 0 {
		if len(starStack) == 0 {
			return false
		}

		if leftStack[len(leftStack)-1] > starStack[len(starStack)-1] {
			return false
		}
		leftStack = leftStack[:len(leftStack)-1]
		starStack = starStack[:len(starStack)-1]
	}

	return true
}
