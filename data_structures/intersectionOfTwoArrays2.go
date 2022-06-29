// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.

package datastructures

func intersect(nums1 []int, nums2 []int) []int {
	counterTable := make(map[int]int)

	for _, val := range nums1 {
		if _, ok := counterTable[val]; !ok {
			counterTable[val] = 0
		}

		counterTable[val]++
	}

	res := make([]int, 0)
	for _, val := range nums2 {
		if counter, ok := counterTable[val]; !ok || ok && counter == 0 {
			continue
		}

		counterTable[val]--
		res = append(res, val)
	}

	return res
}
