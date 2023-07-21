package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}

func strStr(haystack string, needle string) int {
	res := false
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			for j := 0; j < len(needle); j++ {
				if i+j < len(haystack) && haystack[i+j] == needle[j] {
					res = true
				} else {
					res = false
					break
				}
			}
		}
		if res {
			return i
		}
	}
	return -1
}
