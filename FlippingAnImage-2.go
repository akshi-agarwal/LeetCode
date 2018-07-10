
// Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.
// To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].
// To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].
// Notes:
// 1 <= A.length = A[0].length <= 20
// 0 <= A[i][j] <= 1

//In this version of the problem I used bit masking to invert the digits in the matrix
func flipAndInvertImage(A [][]int) [][]int {
	//Reversing  it
	for i := 0; i < len(A); i++ {
		bi := len(A)
		for j := 0; j < (len(A)+1)/2; j++ {
			bi--
			temp := A[i][bi] ^ 1
			A[i][bi] = A[i][j] ^ 1
			A[i][j] = temp

		}
	}
	return A
}

//Time: O(n)
//Here n is the number of elements. No element gets iterated over twice.
