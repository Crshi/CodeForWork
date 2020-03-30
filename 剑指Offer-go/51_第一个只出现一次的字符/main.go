package main

func firstUniqChar(s string) byte {
	hash := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if hash[s[i]] == 1 {
			return s[i]
		}
	}

	var res byte
	res = ' '
	return res
}

func main() {

}
