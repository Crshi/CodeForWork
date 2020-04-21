package main

func numSteps(s string) int {
	count := 0
	for len(s) != 1 {
		count++
		if s[len(s)-1] == 49 {
			tmp := len(s) - 1
			i := tmp
			for i >= 0 && s[i] == 49 {
				i--
			}

			if i == -1 {
				s = "1"
			} else {
				s = s[:i] + "1"
			}

			for i < tmp {
				s += "0"
				i++
			}
		} else {
			s = s[:len(s)-1]
		}
	}

	return count
}

func main() {
	numSteps("111111001110101011")
}
