package editdistance

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected int
	}{
		{"Example1", "horse", "ros", 3},
		{"Example2", "intention", "execution", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.expected {
				t.Errorf("minDistance(%s, %s) = %d; expected %d", tt.word1, tt.word2, got, tt.expected)
			}
		})
	}

}
