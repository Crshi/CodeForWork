package main

type node struct {
	x int
	y int
}
func Sum(x, y int) int {
	sum := 0
	for x != 0 {
		sum += x % 10
		x /= 10
	}
	for y != 0 {
		sum += y % 10
		y /= 10
	}
	return sum
}

func movingCountBFS(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	vis := make([][]bool, m + 1)
	for i := 0; i < m + 1; i ++ {
		vis[i] = make([]bool, n + 1)
	}
	que := []node{node{
		x:0,
		y:0,
	}}
	var ans int = 0
	dir := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(que) != 0 {
		front := que[0]
		que = que[1:]

		if vis[front.x][front.y] {
			continue
		}
		if Sum(front.x, front.y) > k {
			continue
		}
		vis[front.x][front.y] = true
		ans ++
		for i := 0; i < 4; i ++ {
			nx := front.x + dir[i][0]
			ny := front.y + dir[i][1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				que = append(que, node{x:nx, y:ny,})
			}
		}
	}
	return ans
}

func main() {
	println(movingCountBFS(2,3,1))
}
