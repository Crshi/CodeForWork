package main

func spiralOrder(matrix [][]int) []int {
	datas := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			datas = append(datas, matrix[i][j])
		}
	}

	return datas
}

func main() {

}
