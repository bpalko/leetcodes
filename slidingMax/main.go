package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(slidingWindowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) // Output: [3 3 5 5 6 7]
}

// Algorithm:
// Left and right pointers are used to maintain the current window of size k.
// The function returns a slice containing the maximum values of each sliding window of size k.

// Edge cases:
// input is empty
// k is greater than the length of the input slice
// k is less than or equal to 0
// no maximum value exists in the input slice

// slidingWindowMax finds the maximum values in each sliding window of size k.
//
// Parameters:
// nums: The input slice of integers.
// k: The size of the sliding window.
//
// Returns:
// A slice containing the maximum values of each sliding window of size k.
func slidingWindowMax(nums []int, k int) []int {
	// Handle edge cases
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return []int{}
	}

	result := []int{}
	deque := []int{} // Stores indices of elements in the current window

	for i := 0; i < len(nums); i++ {
		// Remove indices that are out of the bounds of the current window
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Remove indices of elements smaller than the current element
		// as they will not be the maximum in the current window
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Add the current index to the deque
		deque = append(deque, i)

		// Append the maximum value to the result once the first window is formed
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
