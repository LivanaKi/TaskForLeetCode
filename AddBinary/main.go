package main

import(
	"fmt"
)

func main(){
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		 a, b = b, a
	}
	result := make([]byte, len(a))
	bIndex := len(b)-1
	var sum, shf byte
	for i := len(a) -1; i >= 0; i--{
		sum = a[i] + shf
		if bIndex >= 0{
		sum += b[bIndex]
		bIndex--
		}
		result[i] = sum % 2 + '0'
		shf = sum >> 1 % 2
	}
	if shf == 0{
		return string(result)
	}
	return "1" + string(result)
   }