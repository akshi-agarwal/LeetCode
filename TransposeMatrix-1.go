// Given a matrix A, return the transpose of A.
// The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.
// Example 1:
// Input: [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[1,4,7],[2,5,8],[3,6,9]]

func transpose(A [][]int) [][]int {
	arr := make([][]int, len(A[0]))
	for i := range arr {
		arr[i] = make([]int, len(A))
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			arr[j][i] = A[i][j]
		}
	}
	return arr
}

//Time: O(n)
