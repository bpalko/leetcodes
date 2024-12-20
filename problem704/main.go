package main

import "fmt"


func main() {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{}, 5, -1},
		{[]int{1, 3}, 3, 1},
	}

	for i, tc := range testCases {
		result := binarySearch(tc.nums, 0, len(tc.nums)-1, tc.target)
		passed := result == tc.expected
		fmt.Printf("Test case %d: %v\n", i+1, passed)
		fmt.Printf("Input: nums=%v, target=%d\n", tc.nums, tc.target)
		fmt.Printf("Expected: %d, Got: %d\n\n", tc.expected, result)
	}
}

func binarySearch(nums []int, lower, upper, target int) int {
	mid := (upper + lower) / 2
	fmt.Printf("mid: %v\n", mid)
	fmt.Printf("lower: %v\n", lower)
	fmt.Printf("upper: %v\n", upper)
	if lower > upper {
		return -1
	}
	if nums[mid] == target {
		fmt.Printf("%v is equal to target\n", nums[mid])
		return mid
	}
	if nums[mid] < target {
		fmt.Printf("%v is less than target\n", nums[mid])
		fmt.Printf("nums[mid+1:]: %v\n", nums[mid+1:])
		lower = mid+1
		return binarySearch(nums, lower, upper, target)
	}
	if nums[mid] > target {
		fmt.Printf("%v is more than target", nums[mid])
		fmt.Printf("nums[:mid-1]: %v\n", nums[:mid-1])
		upper = mid-1
		return binarySearch(nums, lower, upper, target)
	}
	return -1
}


/*
[-1, 0, 3, 5, 9, 12] targ=9

Upper and lower bounds (6 + 0) / 2 = 3
midpoint = 3
check if midpoint[3] = target, else
if midpoint[3] < target
remainder of the list from midpoint+1 == lower bound, upper bound 12

(12 + 9) / = 10

check if midpoint = target, else
if midpoint < target
remainder of the list from midpoint - 1

*/