package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
}

func searchInsert(nums []int, target int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	if res == 0 {
		nums = append(nums, target)
		sort.Ints(nums)
		for j := 0; j < len(nums); j++ {
			if nums[j] == target {
				res = j
			}
		}
	}

	return res
}
