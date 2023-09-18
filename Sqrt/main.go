package main

import(
	"fmt"
	"math"
)

func main(){
	num := 8
	fmt.Println(mySqrt(num))
}

func mySqrt(x int) int {
    return int(math.Sqrt(float64(x)))
}