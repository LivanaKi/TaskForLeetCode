package main

import(
	"fmt"
)

func main(){
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

func removeDuplicates(nums []int) int {
    i := 0
    for j := 1; j < len(nums); j++{
        if nums[i] != nums[j]{
            i++
            nums[i] = nums[j]
        }
    }
    return i+1
}