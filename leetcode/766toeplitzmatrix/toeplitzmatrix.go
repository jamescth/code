package main

func isToeplitzMatrix(matrix [][]int) bool {
	row := len(matrix)
	if row == 1 {
		return true
	}

	column := len(matrix[0])
	for i := 0; i < row-1; i++ {
		for j := 0; j < column-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}

	}
	return true

}
