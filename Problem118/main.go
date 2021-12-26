package main

import (
	"fmt"
	"time"
)
/*
https://leetcode.com/problems/pascals-triangle/
118. Pascal's Triangle
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
 */
func main() {
	timeStart := time.Now()

	fmt.Println(time.Since(timeStart))
}
func generate(numRows int) [][]int {
	out := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		//temp := make([]int,i+1)
		var temp []int
		temp = append(temp,1)
		for j := 0; j < len(out[i-1])-1; j++ {
			temp = append(temp,out[i-1][j]+out[i-1][j+1])
		}
		temp = append(temp,1)
		out = append(out, temp)
	}
	return out
}