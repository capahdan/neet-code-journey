package nqueens

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n    int
		want [][]string
	}{
		{4, [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("TestSolveNQueens(%d)", test.n), func(t *testing.T) {
			got := solveNQueens(test.n)
			assert.Equal(t, test.want, got, "expected %v, got %v", test.want, got)
		})
	}
}
