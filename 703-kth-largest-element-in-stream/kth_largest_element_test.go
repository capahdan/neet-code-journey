package kthlargestelementinstream

import "testing"

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name string
		k    int
		nums []int
		val  int
		want int
	}{
		{name: "test 1", k: 3, nums: []int{4, 5, 8, 2}, val: 3, want: 4},
		{name: "test 2", k: 1, nums: []int{}, val: 5, want: 5},
		{name: "test 4", k: 2, nums: []int{5, -1}, val: 1, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			kl := Constructor(test.k, test.nums)
			got := kl.Add(test.val)
			if got != test.want {
				t.Errorf("KthLargest(%v, %d) = %d, want %d", test.nums, test.k, got, test.want)
			}
		})
	}
}
