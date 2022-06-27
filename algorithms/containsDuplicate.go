// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: true
// Example 2:

// Input: nums = [1,2,3,4]
// Output: false
// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

package algorithms

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, el := range nums {
		if _, ok := set[el]; ok {
			return true
		}
		set[el] = struct{}{}
	}

	return false
}
