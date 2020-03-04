package main

import "strings"

func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func main() {
	println(replaceSpace("We are happy."))
}
