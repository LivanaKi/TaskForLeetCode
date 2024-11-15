package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			ans = append(ans, dontKnow(nums, i)...)
		}
	}

	mp := make(map[[3]int]int)
	for i := 0; i < len(ans); i++ {
		var abc []int
		abc = ans[i]
		mp[[3]int(abc)]++
	}
	var answer [][]int
	for val := range mp {
		answer = append(answer, val[0:3])
	}
	return answer
}

func dontKnow(arr []int, j int) [][]int {

	var ans [][]int
	sm := 0 - arr[j]
	low := j + 1
	high := len(arr) - 1
	for low < high {
		if arr[low] + arr[high] == sm {
			ans = append(ans, []int{arr[j], arr[low], arr[high]})

			if low < high && arr[low] == arr[low+1] {
				low++
			}
			if low < high && arr[high] == arr[high-1] {
				high--
			}
			low++
			high--
		}else if arr[low]+arr[high] > sm {
			high--
		} else {
			low++
		}
	}

	return ans
}

func main() {
	fmt.Println(threeSum([]int{3, -2, 1, 0}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
