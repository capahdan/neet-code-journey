package happynumber

import "testing"

func TestHappyNumber(t *testing.T) {
	testTable := []struct {
		name   string
		n      int
		expect bool
	}{
		{
			name:   "Happy Number",
			n:      19,
			expect: true,
		},
		{
			name:   "Not Happy Number",
			n:      2,
			expect: false,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := IsHappy(test.n)
			if result != test.expect {
				t.Errorf("IsHappy(%d) = %v, expected %v", test.n, result, test.expect)
			}
		})
	}
}
