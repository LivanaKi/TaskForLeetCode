package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	nearest := 100000
	closestSum := 0

	for i := 0; i < len(nums)-2; i++ {
		low := i + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			absSum := int(math.Abs(float64(target - sum)))
			if sum == target {
				return sum
			}
			fmt.Println(sum)
			if absSum < nearest {
				nearest = absSum
				closestSum = sum
			}

			if nums[i]+nums[low]+nums[high] > target {
				high--
			} else {
				low++
			}
		}

	}

	return closestSum
}

func main() {
	fmt.Println(threeSumClosest([]int{0, 3, 97, 102, 200}, 300))
}
