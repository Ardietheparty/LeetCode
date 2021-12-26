package main

import (
	"fmt"
	"time"
)
/*
42. Trapping Rain Water
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it can trap after raining.
 */
func main() {
	timeStart := time.Now()
	//height = [0,1,0,2,1,0,1,3,2,1,2,1]
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(time.Since(timeStart))
}
func trap(height []int) int {
	//contain := make([]int, len(height))
	cont := 0
	temp :=0
	for i := 0; i < len(height)-1; i++ {
		if temp < i {
			temp = i
		}
		for j := i; j < len(height); j++ {
			if temp < height[j] {
				cont++
			}
		}
	}
	return cont
}