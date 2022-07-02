// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
//
// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the previous row.
//
// Example 1:
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true
//
// Example 2:
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false

package datastructures

// binary search approach. O(log(m*n))
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	begin, end := 0, m*n
	for begin < end {
		mid := (end + begin) / 2
		i, j := mid/n, mid%n

		if matrix[i][j] == target {
			return true
		} else if target < matrix[i][j] {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return false
}

// two pointers approach. O(m*log(n))
func searchMatrixTwoPointers(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := 0
	for i < m {
		if matrix[i][n-1] == target {
			return true
		} else if target < matrix[i][n-1] {
			return binarySearch(matrix[i], target)
		} else {
			i++
		}
	}
	return false
}

func binarySearch(row []int, target int) bool {
	begin, end := 0, len(row)
	for begin < end {
		mid := (end + begin) / 2
		if row[mid] == target {
			return true
		} else if target < row[mid] {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return false
}
