package main

import (
	"fmt"
)

func main(){
	nums := []int{4,5,6,7,0,1,2}
	target := 6
	fmt.Println(searc(nums, target))
}

func searc(nums []int, target int) int{
	var result int

	for i:= 0; i < len(nums); i++{
		if target == nums[i]{
			result = i
			break
		}else{
			result = -1
		}
	}
	return result
}