package main

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int, len(deck))
	for _, v := range deck {
		m[v]++
	}
	g := -1
	for _, v := range m {
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g >= 2
}

//求最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1})
}
