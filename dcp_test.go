package dcp

import "testing"

func TestGetMaxCoins(t *testing.T) {

	type TestCase struct {
		name           string
		matrix         [][]int
		expectedResult int
	}

	testCases := []TestCase{
		{
			name: "case 1 - from example",
			matrix: [][]int{
				{0, 3, 1, 1},
				{2, 0, 0, 4},
				{1, 5, 3, 1},
			},
			expectedResult: 12,
		},
		{
			name: "case 2",
			matrix: [][]int{
				{1, 0, 0, 0},
				{0, 1, 6, 0},
				{0, 0, 0, 14},
			},
			expectedResult: 22,
		},
		{
			name: "case 3",
			matrix: [][]int{
				{0, 4, 0, 0},
				{1, 0, 3, 0},
				{2, 0, 20, 0},
			},
			expectedResult: 27,
		},
		{
			name: "case 4",
			matrix: [][]int{
				{1, 3, 1, 1},
				{2, 0, 0, 4},
				{1, 5, 3, 1},
			},
			expectedResult: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getMaxCoins(tc.matrix)

			if result != tc.expectedResult {
				t.Errorf("For matrix %v, expected %d but got %d", tc.matrix, tc.expectedResult, result)
			}
		})
	}
}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix := [][]int{
			{0, 3, 1, 1},
			{2, 0, 0, 4},
			{1, 5, 3, 1},
		}
		getMaxCoins(matrix)
	}
}
