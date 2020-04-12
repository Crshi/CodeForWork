package main

func gameOfLife(board [][]int) {
	datas := make([][]int, 0)

	for i := 0; i < len(board); i++ {
		tmp := make([]int, len(board[i]))
		copy(tmp, board[i])
		datas = append(datas, tmp)
	}

	for i := 0; i < len(datas); i++ {
		for j := 0; j < len(datas[i]); j++ {
			count := 0
			if i-1 >= 0 {
				if j-1 >= 0 && datas[i-1][j-1] == 1 {
					count++
				}

				if datas[i-1][j] == 1 {
					count++
				}

				if j+1 < len(datas[i]) && datas[i-1][j+1] == 1 {
					count++
				}
			}

			if j-1 >= 0 && datas[i][j-1] == 1 {
				count++
			}
			if j+1 < len(datas[i]) && datas[i][j+1] == 1 {
				count++
			}

			if i+1 < len(datas) {
				if j-1 >= 0 && datas[i+1][j-1] == 1 {
					count++
				}

				if datas[i+1][j] == 1 {
					count++
				}

				if j+1 < len(datas[i]) && datas[i+1][j+1] == 1 {
					count++
				}
			}

			if datas[i][j] == 1 {
				if count < 2 || count > 3 {
					board[i][j] = 0
				}
			} else {
				if count == 3 {
					board[i][j] = 1
				}
			}
		}
	}

	return
}

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
}
