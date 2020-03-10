package main

func hammingDistance(x int, y int) int {
	dis := 0
	var sx [32]int
	var sy [32]int
	xi := 0
	yi := 0

	for x > 0 {
		sx[xi] = x % 2
		xi++
		x = x / 2
	}

	for y > 0 {
		sy[yi] = y % 2
		yi++
		y = y / 2
	}

	for i := 0; i < 32; i++ {
		if sx[i] != sy[i] {
			dis++
		}
	}

	return dis
}

func main() {
	println(hammingDistance(1, 4))
}
