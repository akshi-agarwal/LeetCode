func flipAndInvertImage(A [][]int) [][]int {
	//Reversing  it
	for i := 0; i < len(A); i++ {
		bi := len(A)
		for j := 0; j < len(A)/2; j++ {
			bi--
			temp := A[i][bi]
			A[i][bi] = A[i][j]
			A[i][j] = temp
			
		}
	}

	//Inverting it
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
	}
	return A
}

Time: O(n^2)
