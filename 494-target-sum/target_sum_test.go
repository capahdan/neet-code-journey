package targetsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expedted int
	}{
		{
			name:     "test1",
			nums:     []int{1, 1, 1, 1, 1},
			target:   3,
			expedted: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expedted, TargetSum(tt.nums, tt.target))
		})
	}

}
