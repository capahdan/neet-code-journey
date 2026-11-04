package regulerexpressionmatching

import "testing"

func TestRegulerExpression(t *testing.T) {
	testcase := []struct {
		s        string
		p        string
		expected bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
	}

	for _, tc := range testcase {
		result := IsMatch(tc.s, tc.p)
		if result != tc.expected {
			t.Errorf("IsMatch(%s, %s) = %v, expected %v", tc.s, tc.p, result, tc.expected)
		}
	}

}
