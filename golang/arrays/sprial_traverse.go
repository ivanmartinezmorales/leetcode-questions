package main

func SpiralTraverse(array [][]int) []int {
	// Write a function that takes an n x m 2-D array and
	// returns a 1D array of all the arrays elements in spiral order

	// Quick smoketest
	if len(array) == 0 {
		return []int{}
	}

	result := []int{}
	spiralFill(array, 0, len(array)-1, 0, len(array[0])-1, &result)
	return result
}

func spiralFill(array [][]int, startCol, endCol, startRow, endRow int, result *[]int) {

	for col := startCol; col <= endCol; col++ {
		*result = append(*result, array[startRow][col])
	}

	for row := startRow; row <= endRow; row++ {
		*result = append(*result, array[row][endCol])
	}

	for col := endCol - 1; col >= startCol; col-- {
		if startRow == endRow {
			break
		}
		*result = append(*result, array[endRow][col])
	}

	for row := endRow - 1; row >= startRow; row-- {
		if startCol == endCol {
			break
		}
		*result = append(*result, array[row][startCol])
	}

	spiralFill(array, startRow+1, endRow-1, startCol+1, endCol-1, result)
}
