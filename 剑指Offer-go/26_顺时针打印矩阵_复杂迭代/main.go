package main

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return nil
	}

	datas := make([]int, 0)

	minx := 0
	maxx := len(matrix)

	miny := 0
	maxy := len(matrix[0])

	step := 0
	size := len(matrix[0]) * len(matrix)

	for step < size {
		for j := miny; j < maxy && step < size; j++ {
			datas = append(datas, matrix[minx][j])
			step++
		}
		minx++

		for i := minx; i < maxx && step < size; i++ {
			datas = append(datas, matrix[i][maxy-1])
			step++
		}
		maxy--

		for j := maxy - 1; j >= miny && step < size; j-- {
			datas = append(datas, matrix[maxx-1][j])
			step++
		}

		maxx--

		for i := maxx - 1; i >= minx && step < size; i-- {
			datas = append(datas, matrix[i][miny])
			step++
		}

		miny++
	}

	return datas
}

// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]

func main() {

}
