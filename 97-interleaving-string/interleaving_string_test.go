package interleavingstring

import (
	"testing"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		s3       string
		expected bool
	}{
		{
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbcbcac",
			expected: true,
		},
		{
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbbaccc",
			expected: false,
		},
		{
			s1:       "",
			s2:       "",
			s3:       "",
			expected: true,
		},
		{
			s1:       "a",
			s2:       "",
			s3:       "a",
			expected: true,
		},
		{
			s1:       "",
			s2:       "b",
			s3:       "b",
			expected: true,
		},
		{
			s1:       "a",
			s2:       "b",
			s3:       "ab",
			expected: true,
		},
		{
			s1:       "a",
			s2:       "b",
			s3:       "ba",
			expected: true,
		},
		{
			s1:       "a",
			s2:       "b",
			s3:       "ac",
			expected: false,
		},
		{
			s1:       "",
			s2:       "",
			s3:       "a",
			expected: false,
		},
		{
			s1:       "a",
			s2:       "",
			s3:       "",
			expected: false,
		},
		{
			s1:       "a",
			s2:       "b",
			s3:       "",
			expected: false,
		},
		{
			s1:       "a",
			s2:       "b",
			s3:       "abc",
			expected: false,
		},
	}

	for _, test := range tests {
		result := isInterleave(test.s1, test.s2, test.s3)
		if result != test.expected {
			t.Errorf("isInterleave(%s, %s, %s) = %v, expected %v", test.s1, test.s2, test.s3, result, test.expected)
		}
	}
}
