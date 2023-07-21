package main

import(
	"fmt"
)

func main(){
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
    res, maxL := "", 0
    for i := range s {
        l, r := i, i 
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if r - l + 1 > maxL {
                res = s[l:r + 1]
                maxL = r - l + 1
            }
            l--
            r++
        }
        l, r = i, i + 1 
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if r - l + 1 > maxL {
                res = s[l:r + 1]
                maxL = r - l + 1
            }
            l--
            r++
        }
    }
    return res
}