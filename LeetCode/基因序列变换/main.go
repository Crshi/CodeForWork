package main

func minMutation(start string, end string, bank []string) int {
	changeCount := 0

	start1 := []byte(start)
	end1 := []byte(end)

	for i := 0; i < len(end1); i++ {
		if start1[i] != end1[i] {
			start1[i] = end1[i]
			if isExist(bank, string(start1[:])) {
				changeCount++
			} else {
				return -1
			}
		}
	}

	return changeCount
}

func isExist(bank []string, target string) bool {
	for _, x := range bank {
		if x == target {
			return true
		}
	}

	return false
}

func main() {
	data := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}

	println(minMutation("AACCGGTT", "AAACGGTA", data))
}
