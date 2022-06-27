// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Example 2:

// Input: nums = [1]
// Output: 1
// Example 3:

// Input: nums = [5,4,-1,7,8]
// Output: 23

package algorithms

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// dynamic programming approach
func maxSubArrayDP(nums []int) int {
	maxSumEnding := make([]int, len(nums))

	maxSumEnding[0] = nums[0]
	maxRes := maxSumEnding[0]
	for i := 1; i < len(nums); i++ {
		maxSumEnding[i] = maxInt(maxSumEnding[i-1]+nums[i], nums[i])
		maxRes = maxInt(maxRes, maxSumEnding[i])
	}

	return maxRes
}

// Kadane algorithm based on method above
func maxSubArrayKadane(nums []int) int {
	maxSumEndingHere := nums[0]
	maxRes := maxSumEndingHere

	for i := 1; i < len(nums); i++ {
		maxSumEndingHere = maxInt(maxSumEndingHere+nums[i], nums[i])
		maxRes = maxInt(maxRes, maxSumEndingHere)
	}
	return maxRes
}

// TODO: divide and conquer approach
