package main

func sum(m int, n int) int {
	sum := 0
	for m != 0 {
		sum += m % 10
		m = m / 10
	}

	for n != 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}

func movingCountDFS(m int, n int, k int) int {
	data := make([][]bool, m)
	for i := 0; i < m; i++ {
		data[i] = make([]bool, n)
	}

	return dfs(data, 0, 0, k)
}

func dfs(data [][]bool, m int, n int, k int) int {

	if sum(m, n) > k {
		return 0
	}

	if m >= len(data) || m < 0 || n >= len(data[0]) || n < 0 {
		return 0
	}

	if data[m][n] == true {
		return 0
	}

	data[m][n] = true

	sum := 1

	sum += dfs(data, m+1, n, k)
	sum += dfs(data, m-1, n, k)
	sum += dfs(data, m, n+1, k)
	sum += dfs(data, m, n-1, k)

	return sum
}
