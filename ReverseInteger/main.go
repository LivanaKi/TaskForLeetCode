package main

import(
	"fmt"
	"strings"
	"strconv"
	"math"
)

func main(){
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	var result []string
	var res strings.Builder
	var ans int

	s := strconv.Itoa(x)
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, string(s[i]))
	}
	for _, val := range result {
		if val == "-" {
			val = ""
		}
		res.WriteString(val)
	}
	if x < 0 {
		x, _ = strconv.Atoi("-" + res.String())
		ans = x
	} else {
		x, _ = strconv.Atoi(res.String())
		ans = x
	}
	if x >= math.MaxInt32 || x <= math.MinInt32 {
		ans = 0
	}
	return ans
}