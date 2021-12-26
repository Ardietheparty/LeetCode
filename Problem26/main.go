package main

import (
	"fmt"
	"time"
)
/*
26. Remove Duplicates from Sorted Array
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
 */
func main() {
	timeStart := time.Now()
	s := []int{1,1,2}
	fmt.Println(removeDuplicates(s))
	fmt.Println(time.Since(timeStart))
}
func removeDuplicates(nums []int) int {
	l := len(nums)
	i:=0
	n := 0
	q := 0
	hold := make([]int,l)
	for i < l {
		fmt.Println(nums[i])
		hold[n]=nums[i]
		i+=check(nums,i,l)
		q++
		n++
	}
	//fmt.Println(hold)
	nums = nil
	nums = hold
	return l-q
}
func check(nums []int, i int, l int) int {
	hold := nums[i]
	a := 0

	for j:=i; j<l; j++{
		if hold != nums[j] {
			return j-i
		}
		a=j-i
	}
	return a+1
}
