package main

import (
	"fmt"
	"time"
)
/*
https://leetcode.com/problems/pascals-triangle-ii/
119. Pascal's Triangle II
Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
 */
func main() {
	timeStart := time.Now()

	fmt.Println(time.Since(timeStart))
}
func getRow(rowIndex int) []int  {
	out := []int{1}
	for i := 1; i < rowIndex+1; i++ {
		//temp := make([]int,i+1)
		var temp []int
		temp = append(temp,1)
		for j := 0; j < len(out)-1; j++ {
			temp = append(temp,out[j]+out[j+1])
		}
		temp = append(temp,1)
		out = temp
	}
	return out
}