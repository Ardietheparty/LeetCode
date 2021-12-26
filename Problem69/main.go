package main

import (
	"fmt"
	"time"
)
/*
69. Sqrt(x)
Given a non-negative integer x, compute and return the square root of x.

Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.

Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.
 */

func main() {
	timeStart := time.Now()
	fmt.Println(mySqrt(10))
	fmt.Println(time.Since(timeStart))
}
func mySqrt(x int) int {
	ret := x
	for ret*ret > x {
		ret = (ret+x/ret)/2
	}
	return ret
}