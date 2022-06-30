// Given an integer numRows, return the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
//
// Example 1:
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// Example 2:
//
// Input: numRows = 1
// Output: [[1]]

package datastructures

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)

		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

// similar solution as above one with row traversal to the middle (triangle is symmetric)
func generateTraversalToTheMiddle(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)

		triangle[i][0], triangle[i][i] = 1, 1
		mid := i >> 1 // mid := i / 2
		for j := 1; j <= mid; j++ {
			val := triangle[i-1][j-1] + triangle[i-1][j]
			triangle[i][j], triangle[i][i-j] = val, val
		}
	}
	return triangle
}
