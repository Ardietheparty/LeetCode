package main

import (
	"fmt"
	"sort"
	"time"
)
/*
973. K Closest Points to Origin
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).
 */
func main() {
	timeStart := time.Now()
	p:= [][]int{{1,3},{-2,2}}
	fmt.Println(kClosest(p,1))
	fmt.Println(time.Since(timeStart))
}
func kClosest(points [][]int, k int) [][]int {
	var out [][]int
	var held [][]float64
	for i := 0; i < len(points); i++ {
		hold := points[i][0]* points[i][0] + points[i][1]*points[i][1]
		//temp := math.Sqrt(float64(hold))
		//fmt.Println(points[i],temp)
		held = append(held,[]float64{float64(i),float64(hold)})
	}
	//fmt.Println(held)
	sort.SliceStable(held, func(i, j int) bool {
		return held[i][1] < held[j][1]
	})
	//fmt.Println(held)
	for i := 0; i < k; i++ {
		out = append(out,points[int(held[i][0])])
	}
	return out
}