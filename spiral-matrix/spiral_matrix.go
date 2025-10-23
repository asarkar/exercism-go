package spiralmatrix

// `SpiralMatrix` generates a size x size spiral matrix filled with numbers 1..size*size.
func SpiralMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	topRow, leftCol := 0, 0
	bottomRow, rightCol := size-1, size-1
	val := 0
	total := size * size

	for val < total {
		// Fill top row
		for col := leftCol; col <= rightCol; col++ {
			val++
			matrix[topRow][col] = val
		}
		topRow++

		// Fill right column
		for row := topRow; row <= bottomRow; row++ {
			val++
			matrix[row][rightCol] = val
		}
		rightCol--

		// Fill bottom row
		for col := rightCol; col >= leftCol; col-- {
			val++
			matrix[bottomRow][col] = val
		}
		bottomRow--

		// Fill left column
		for row := bottomRow; row >= topRow; row-- {
			val++
			matrix[row][leftCol] = val
		}
		leftCol++
	}

	return matrix
}
