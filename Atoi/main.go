package main

import(
	"fmt"
	"unicode"
	"math"
)

func main(){
	fmt.Println(myAtoi("42"))
}

func myAtoi(s string) int {
	var result int
	
	i := 0

	for i < len(s) && s[i] == ' '{
		i++
	}
	sing  := 1
	if i < len(s) && (s[i] == '+' || s[i] == '-'){
		if s[i] == '-'{
			sing = -1
		}else{
			sing = 1
		}
		i++
	}


	for i < len(s) && unicode.IsNumber(rune(s[i])){
		result = result*10 + int(rune(s[i]) - '0')
		if result > math.MaxInt32 {
			if sing == 1{
			return math.MaxInt32
		}else{
			return math.MinInt32
		}
	}
		i++
	}
return result*sing
}