package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
}

func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	result := false
	for i := 0; i < len(nums); i++ {
		_, ok := mp[nums[i]]
		if ok {
			return true
		}
		mp[nums[i]] = false
	}

	return result
}
