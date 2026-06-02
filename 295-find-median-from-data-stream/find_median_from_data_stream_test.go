package findmedianfromdatastream

import "testing"

func TestFindMedianFromDataStream(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want float64
	}{
		{name: "test 1", nums: []int{1, 2, 3, 4, 5}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			medianFinder := Constructor()
			for _, num := range test.nums {
				medianFinder.AddNum(num)
			}
			got := medianFinder.FindMedian()
			if got != test.want {
				t.Errorf("FindMedianFromDataStream(%v) = %f, want %f", test.nums, got, test.want)
			}
		})
	}
}
