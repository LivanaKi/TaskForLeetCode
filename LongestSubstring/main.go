package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
func lengthOfLongestSubstring(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		set := make(map[byte]bool)
		length := 0
		for j := i; j < len(s); j++ {
			if set[s[j]] {
				break
			}

			set[s[j]] = true
			length++

			fmt.Println(set)
		}

		if result < length {
			result = length
		}
	}

	return result
}
