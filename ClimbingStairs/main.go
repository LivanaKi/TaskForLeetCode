package main

import (
	"fmt"
)

func main(){
	n := 2
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	var res int
	one := 1
    two := 1
    for i := 2; i <= n; i++ {
        res = one + two
        two = one
        one = res
    }
    return res
}