package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(longestCommonPrefix([]string{"aaa","aa","aaa"}))
}

func checkPrefix(strs []string, str string) bool{
    var s bool

    for i := 0; i < len(strs); i++{
        s = strings.HasPrefix(strs[i], str)
        if s == false{
            break
        }
    }
    return s
}

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
        return ""
    }

    if len(strs) == 1{
        return strs[0]
    }

    str := strs[0]

    var s string

    for i := 0; i <= len(str); i++{
        if checkPrefix(strs, str[:i]){
           s = str[:i]
        }else{
            break
        }
    }
    return s
}