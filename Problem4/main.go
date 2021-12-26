package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	timeStart := time.Now()
	a := []int{9,14,16,17}
	b := []int{1,4,6}
	ab := append(a, b...)

	sort.Ints(ab)
	fmt.Println(med(ab))
	//fmt.Println(ab)
	//fmt.Println(merg(a,b,5))
	fmt.Println(findMedianSortedArrays(a,b))

	fmt.Println(time.Since(timeStart))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var out float64
	a := len(nums1)
	b := len(nums2)
	c := a+b
	even := true
	if c%2 == 0 {
		c=c/2
		c++
		even = true
	} else {
		c=c+1
		c=c/2
		even = false
	}
	lel := merg(nums1,nums2,c)
	if even {
		out = float64(lel[c-1]+lel[c-2])
		out = out/2
	} else {
		out = float64(lel[c-1])
	}
	return out
}
func med(num []int) float64 {
	N := len(num)
	n := 0.0
	if N%2 != 0 {
		n = float64(num[N/2])
	} else {
		a := (N-1) / 2
		n =  float64(num[a]+num[a+1])
		n = n / 2
	}
	return n
}
func meds(num []int) []int {
	N := len(num)
	var out []int
	if N%2 != 0 {
		out = append(out,num[N/2])
	} else {
		a := (N-1) / 2
		out = append(out,num[a])
		out = append(out,num[a+1])
	}
	return out
}
func merg(a []int, b []int, c int) []int {
	A := len(a)
	B := len(b)
	C := c
	j,k := 0,0
	var out []int
	for i := 0; i < C; i++ {
		if j >= A {
			out = append(out, b[k])
			k++
		} else if k >= B {
			out = append(out,a[j])
			j++
		} else {
			m := a[j]
			n := b[k]
			if m < n {
				out = append(out,m)
				j++
			} else {
				out = append(out,n)
				k++
			}
		}
	}
	return out
}