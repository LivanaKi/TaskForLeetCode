package main

import (
	"fmt"
	"strconv"
	"strings"
	//	"strings"
)

func letterCombinations(digits string) []string {
	var ans []string

	mapWord := map[int][]string{
		0: nil,
		1: nil,
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	if len(digits) >= 4 {
		fmt.Println("--Error-- (length digit more than 4)")
		return nil
	}
	if len(digits) < 1 {
		fmt.Println("Can you enter data")
		return nil
	}
	if len(digits) == 1 {
		num, err := strconv.Atoi(digits)

		if err != nil {
			fmt.Println("It's not number")
			fmt.Println(err)
			return nil
		}
		return mapWord[num]
	}
	if len(digits) > 1{
		ans = splitLine(digits, mapWord)
	}
	return ans
}

func splitLine(digit string, mapWord map[int][]string) []string{
	result := []string{""}
	for _, val := range digit{
		num, err := strconv.Atoi(string(val))
		if err != nil{
			fmt.Println(err)
		}
		var temprory []string
		 temprory = append(temprory, mapWord[num]...)
		 result = joinLen(result, temprory)
	}
	return result
}

func joinLen(result, temprory []string) []string{
	var res []string
	for i := 0; i < len(result); i++{
		for j := 0; j < len(temprory); j++{
			res = append(res, strings.Join([]string{result[i], temprory[j]}, ""))
		}
	}
	return res
}

func main() {
	fmt.Println(letterCombinations("235"))
}
