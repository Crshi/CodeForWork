package main

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0 ; i < len(grid) ; i++ {
		for j := 0 ; j < len(grid[0]) ;j ++ {
			tmp := getArea(grid,i,j)
			if max < tmp {
				max = tmp
			}
		}
	}

	return max
}

func getArea(grid [][]int,i int,j int) int {
	area := 0
	if grid[i][j] == 1 {
		grid[i][j] = 0
		area = 1
		if i-1 >= 0 {
			area += getArea(grid,i-1,j)
		}
		if i+1 <len(grid) {
			area += getArea(grid,i+1,j)
		}
		if j-1 >= 0 {
			area += getArea(grid,i,j-1)
		}
		if j+1 <len(grid[0]) {
			area += getArea(grid,i,j+1)
		}
		return area
	} else {
		return area
	}
}

func main() {

}
