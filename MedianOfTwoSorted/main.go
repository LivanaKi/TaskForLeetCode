package main

import( 
	"fmt"
	"sort"
)

func main(){
	nums1 := []int{1, 3}
	nums2 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var num []int
	var res float64
    num = append(num, nums1...)
	num = append(num, nums2...)

	sort.Ints(num)

	if len(num)%2 == 0{
		index := len(num)/2
		index2 := index-1
		res = float64(num[index2] + num[index])/2

	}else{
		index := len(num)/2
		res = float64(num[index])
	}
	return res
}