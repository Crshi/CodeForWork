package main

//DP
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}

			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}

			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	grid := [][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}

	maxValue(grid)
}
