package main

import (
	"fmt"
	"time"
)
/*
50. Pow(x, n)
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
 */
func main() {
	timeStart := time.Now()
	fmt.Println(myPow(2,0))
	fmt.Println(myPow(2,1))
	fmt.Println(myPow(2,2))
	fmt.Println(myPow(2,3))
	fmt.Println(myPow(2,4))
	fmt.Println(myPow(2,-1))
	fmt.Println(myPow(2,-2))
	fmt.Println(myPow(2,-3))

	fmt.Println(time.Since(timeStart))
}
func myPow(x float64, n int) float64 {
	neg := false
	a:= x
	even  := n%2==0
	if x == 1 {
		return 1
	}
	if x == -1 {
		if even {
			return 1
		}
		return -1
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		neg = true
		n = -n
	}
	for i := 1; i < n; i++ {
		x *= a
	}
	if neg {
		return 1/x
	}
	return x
}