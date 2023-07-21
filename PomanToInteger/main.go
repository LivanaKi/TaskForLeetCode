package main

import(
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	sum := 0

	literal := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for i, _ := range s {
		if i != len(s)-1 {
			num := literal[string(s[i])]
			next := literal[string(s[i+1])]
			if next > num {
				sum -= num
			} else {
				sum += num
			}

		} else {
			sum += literal[string(s[i])]
		}
	}
	return sum
}
