package main

import (
	"fmt"
	"time"
)
/*
9. Palindrome Number
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

    For example, 121 is a palindrome while 123 is not.

 */
func main() {
	timeStart := time.Now()
	fmt.Println(isPalindrome(11211))

	fmt.Println(time.Since(timeStart))
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	d := digi(x)
	l := len(d)
	for i := 0; i < l/2; i++ {
		if d[i] != d[l-i-1] {
			return false
		}
	}
	return true
}
func digi(x int) []int {
	var out []int
	for x != 0 {
		out = append(out,x%10)
		x = x/10
	}
	return out
}