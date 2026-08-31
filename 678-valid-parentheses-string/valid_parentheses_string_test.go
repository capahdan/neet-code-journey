package validparenthesesstring

import "testing"

func TestIsValidString(t *testing.T) {

	tests := []struct {
		input    string
		expected bool
	}{
		// {"()", true},
		// {"(*)", true},
		// {"(*))", true},
		// {"((", false},
		{"*((", false},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := IsValidString(tt.input); got != tt.expected {
				t.Errorf("IsValidString(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}

}
