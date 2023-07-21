package main

import(
	"fmt"
	"strings"
)

func main(){
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
    var sl []string
	sl = strings.Fields(s)
	return len(sl[len(sl)-1])
}