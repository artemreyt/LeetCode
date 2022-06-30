// In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.
// You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.
// The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.
// If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.
//
// Example 1:
// Input: mat = [[1,2],[3,4]], r = 1, c = 4
// Output: [[1,2,3,4]]

package datastructures

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c || m == r && n == c {
		return mat
	}

	resMat := make([][]int, 0, r)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i*n+j)%c == 0 {
				resMat = append(resMat, make([]int, 0, c))
			}
			resMat[len(resMat)-1] = append(resMat[len(resMat)-1], mat[i][j])
		}
	}
	return resMat
}
