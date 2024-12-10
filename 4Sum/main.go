package main

import (
	"fmt"
	"sort"
)

func fourSum1(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	length2 := len(nums) - 2
	length3 := length2 - 1
	for i := 0; i < length3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < length2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			left, right := j+1, len(nums) -1
			for left < right {
				r := nums[i] + nums[j] + nums[left] + nums[right]
				if r < target {
					left++
				} else if r > target {
					right--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right {
						if nums[left] == nums[left+1] {
							left++
						} else if nums[right] == nums[right-1] {
							right--
						} else {
							break
						}
					}
					left++
					right--
				}
			}
		}
	}
	return result
}


func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			ans = append(ans, dontKnow(nums, i, j, target)...)
		}
	}
	return ans
}

func dontKnow(nums []int, i, j, target int) [][]int {

	var result [][]int
	left, right := j+1, len(nums) -1
			for left < right {
				r := nums[i] + nums[j] + nums[left] + nums[right]
				if r < target {
					left++
				} else if r > target {
					right--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right {
						if nums[left] == nums[left+1] {
							left++
						} else if nums[right] == nums[right-1] {
							right--
						} else {
							break
						}
					}
					left++
					right--
				}
			}
	return result
}

func main() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{1,0,-1,0,-2,2}, 0))
	fmt.Println(fourSum([]int{4,-9,-2,-2,-7,9,9,5,10,-10,4,5,2,-4,-2,4,-9,5}, -13))
}
