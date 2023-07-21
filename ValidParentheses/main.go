package main

import(
	"fmt"
)

func main(){
	fmt.Println(isValid("(]"))
}
func isValid(s string) bool {
    stack := make([]rune, 0)
   
    list := map[rune]rune {
        '{': '}',
        '[': ']',
        '(': ')',
    }
    
    for _, val := range s{
    if sym, ok := list[val]; ok{
        stack = append(stack, sym)
    } else{
        if len(stack) == 0 || stack[len(stack)-1] != val{
            return false
        }
        stack = stack[:len(stack)-1]
    }
    }
    return len(stack) == 0   
   }
