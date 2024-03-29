package main

import(
	"fmt"
)

func main(){
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    ans := len(strs[0])
    for {
        if ans == 0 {
            return ""
        }
        prefix := strs[0][:ans]
        flag := true
        for i := 1; i < len(strs); i++ {
            if len(strs[i]) < ans || strs[i][:ans] != prefix {
                flag = false
                break
            }
        }
        if flag {
            return prefix
        } else {
            ans--
        }
    }
}