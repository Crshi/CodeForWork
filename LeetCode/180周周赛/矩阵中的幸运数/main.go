package main

func luckyNumbers (matrix [3][3]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	tag := make([][]bool,m)
	for i := 0 ;i < m ;i ++ {
		tag[i] = make([]bool,n)
	}
	t := make([]int,0)

	for i := 0; i< m ;i ++ {
		for j := 0; j < n ; j ++ {
			if tag[i][j] == true {
				continue
			}

			if validate(matrix,tag,i,j,m,n) {
				t = append(t,matrix[i][j])
			}
		}
	}

	return t
}

func validate(matrix [3][3]int,tag [][]bool,i int,j int,m int,n int) bool {
	for t := 0;t <m;t++ {
		if t == i {
			continue
		}

		if matrix[t][j] > matrix[i][j]{
			return false
		}
	}

	for t := 0;t <n;t++ {

		if t == j {
			continue
		}

		if matrix[i][t] < matrix[i][j]{
			return false
		}
	}

	for t := 0;t <m ;t++ {
		tag[t][j] = true
	}

	for t := 0;t <n;t++ {
		tag[i][t] = true
	}

	return true
}

//在同一行的所有元素中最小
//在同一列的所有元素中最大
func main() {
	matrix := [3][3]int{{3,7,8},{9,11,13},{15,16,17}}
	t := luckyNumbers(matrix)
	println(t)
}
