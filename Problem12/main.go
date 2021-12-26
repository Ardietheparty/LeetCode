package main

import (
	"fmt"
	"time"
)

/*
12. Integer to Roman

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
 */
func main() {
	timeStart := time.Now()
	fmt.Println(intToRoman(1994))

	fmt.Println(time.Since(timeStart))
}
func intToRoman(num int) string {
	str := ""
	for num != 0 {
		if num >= 1000 {
			str += "M"
			num -= 1000
		} else if num >= 500 {
			str +="D"
			num -= 500
		} else if num >= 100 {
			str += "C"
			num -= 100
		} else if num >= 50 {
			str += "L"
			num -= 50
		} else if num >= 10 {
			str += "X"
			num -= 10
		} else if num >= 5 {
			str += "V"
			num -= 5
		} else if num >= 1 {
			str += "I"
			num -= 1
		}
	}
	return str
}