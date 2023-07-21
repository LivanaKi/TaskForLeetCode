package main

import "fmt"

func main(){
	var nums []int
	var target int

	nums = []int{2, 7, 11, 15}
	target = 9

	fmt.Println(twoSum(nums, target))

}
func twoSum(nums []int, target int) []int {
    var result []int

    for i := 0; i < len(nums); i++{
        for j := 1; j < len(nums); j++{
            res := nums[i] + nums[j]
            if res == target && i < j{
                result = append(result, i, j)
            }
        }
    }
    return result
}