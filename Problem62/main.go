package main

import (
	"fmt"
	"time"
)
/*
Unique Paths
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?
 */

func main() {
	timeStart := time.Now()

	fmt.Println(time.Since(timeStart))
}
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	m--
	n--
	if m < n {
		m = m + n
		n = m - n
		m = m - n
	}

	res, j:=1, 1
	for i:= m +1; i<=m+n; i++ {
		res *=i
		res /=j
		j++
	}
	return res

}
