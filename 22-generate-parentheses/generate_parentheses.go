package generateparentheses

func generateParenthesis(n int) []string {
	result := stack{}
	result.n = n

	result.dfs(0, 0, "")

	return result.res
}

type stack struct {
	res []string
	n   int
}

func (s *stack) dfs(l, r int, str string) {
	if s.n*2 == len(str) {
		s.res = append(s.res, str)
		return
	}
	if l < s.n {
		s.dfs(l+1, r, str+"(")
	}
	if r < l {
		s.dfs(l, r+1, str+")")
	}
}
