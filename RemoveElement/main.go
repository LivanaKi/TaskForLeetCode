package main

import(
	"fmt"
)

func main(){
	fmt.Println(removeElement([]int{3,2,2,3}, 3))
}

func removeElement(nums []int, val int) int {
	count := 0
	   
	   for i:=0; i<len(nums); i++ {
		   nums[i-count] = nums[i]
		   if nums[i]==val {
			   count++
		   }
	   }
	   return len(nums)-count
   }