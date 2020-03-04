package main

func findNumberIn2DArray(matrix [5][5]int, target int) bool {
	i := len(matrix) -1
	j := 0

	for i > -1 && j < len(matrix[i]){
		if matrix[i][j] == target{
			return true
		}

		if matrix[i][j] > target{
			i--
		}else{
			j++
		}
	}

	return false
}

func main() {
	var matrix =
		[5][5]int {{1,   4,  7, 11, 15}, {2,   5,  8, 12, 19}, {3,   6,  9, 16, 22}, {10, 13, 14, 17, 24},	{18, 21, 23, 26, 30}}

	println(findNumberIn2DArray(matrix,5))
}
