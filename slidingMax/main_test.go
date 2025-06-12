package main

import "testing"

func TestSlidingWindowMax(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Empty input",
			nums:     []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "k greater than length",
			nums:     []int{1, 2},
			k:        3,
			expected: []int{},
		},
		{
			name:     "k less than or equal to 0",
			nums:     []int{1, 2, 3},
			k:        0,
			expected: []int{},
		},
		{
			name:     "Valid input",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "Single element window",
			nums:     []int{1, 2, 3},
			k:        1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Window size equals array length",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: []int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slidingWindowMax(tt.nums, tt.k)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %v, got %v", tt.expected, result)
					break
				}
			}
		})
	}
}
