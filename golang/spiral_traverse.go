package main

func SprialTraverse(array [][]int) []int {
		// if the matrix is empty, return an empty array
		if len(array) == 0 {
				return []int{}
		}

		result := []int{}
		startRow, endRow := 0, len(array) - 1
		startCol, endCol := 0, len(array[0]) - 1

		for startRow <= endRow && startCol <= endCol {
				for col := startCol; col <= endCol; col++ {
						result = append(result, array[startRow][col])
				}

				for row := startRow; row <= endRow; row++ {
						result = append(result, array[row][endCol])
				}
		}

		return result
}