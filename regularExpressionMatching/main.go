package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return false
	}

	first := len(s) > 0 && (s[0] == p[0] || p[0] == '.')


	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first && isMatch(s[1:], p))
	} else {
		return first && isMatch(s[1:], p[1:])
	}
}

func main() {
	fmt.Println(isMatch("aa", "a"))
}
