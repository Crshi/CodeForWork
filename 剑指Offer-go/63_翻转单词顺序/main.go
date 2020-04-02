package main

import "strings"

func reverseWords(s string) string {
	devides := strings.Split(s, " ")
	t := ""
	for i := len(devides) - 1; i >= 0; i-- {
		if devides[i] == "" {
			continue
		}

		t += devides[i] + " "
	}

	i := len(t) - 1
	for i >= 0 {
		if t[i] == ' ' {
			t = t[:len(t)-1]
		} else {
			break
		}
		i--
	}

	return t
}

func main() {
	reverseWords("           ")
}
