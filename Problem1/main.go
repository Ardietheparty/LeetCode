package main
/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
 */
import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()


	fmt.Println(time.Since(timeStart))
}
func twoSums(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			h := nums[i]+nums[j]
			if h == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
