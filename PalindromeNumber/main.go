package main

import(
	"fmt"
)

func main(){
	fmt.Println(isPalindrome(121))
}
func isPalindrome(x int) bool {
	var y int
	var z int
	g := x

	if x < 0 {
		return false
	}
	for y < g {
		y := g % 10
		g /= 10
		z = z*10 + y
	}
	if x == z {
		return true
	} else {
		return false
	}
}