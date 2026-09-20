package surroundedregion

import "testing"

func TestSurrounedRegion(t *testing.T) {
	testTable := []struct {
		input    [][]byte
		expected [][]byte
	}{
		{
			input:    [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			expected: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
	}

	for _, tc := range testTable {
		solve(tc.input)
		if !compareBoards(tc.input, tc.expected) {
			t.Errorf("SurrounedRegion(%v) = %v, want %v", tc.input, tc.input, tc.expected)
		}
	}
}

func compareBoards(board1, board2 [][]byte) bool {
	if len(board1) != len(board2) || len(board1[0]) != len(board2[0]) {
		return false
	}

	for i := 0; i < len(board1); i++ {
		for j := 0; j < len(board1[0]); j++ {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}

	return true
}
