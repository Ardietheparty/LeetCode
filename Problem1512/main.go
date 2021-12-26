package main

import (
	"fmt"
	"time"
)
/*
1512. Number of Good Pairs
Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.
 */
func main() {
	timeStart := time.Now()

	fmt.Println(time.Since(timeStart))
}
func numIdenticalPairs(nums []int) int {
	l := len(nums)
	pairs := 0
	for i := 0; i < l; i++ {
		for j:=i+1; j<l; j++ {
			if nums[i] == nums[j] {
				pairs++
			}
		}
	}
	return pairs
}