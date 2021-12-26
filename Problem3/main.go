package main

import (
	"fmt"
	"strings"
	"time"
)
/*
3. Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without repeating characters.
 */
func main() {
	timeStart := time.Now()
	s := "abcabcbb"
	fmt.Println(len(strings.Split(" ","")))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("abbccddeeffgg"))
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))

	fmt.Println(time.Since(timeStart))
}
/*
func lengthOfLongestSubstring(s string) int {
	r := strings.Split(s,"")
	length := len(r)
	lhold := 0
	var temp []string
	for i := 0; i < length; i++ {
		temp = nil
		a := i
		b := i+1
		check := false
		if b < length {
			check = true
		}
		temp = append(temp, r[a])
		for check {
			re := len(temp)
			tes := true
			if b+1 < length {
				check = true
			} else {
				check = false
			}
			for q := 0; q < re; q++ {
				if temp[q] == r[b] {
					tes = false
					check = false
				}
			}
			if tes {
				temp = append(temp,r[b])
			}
			b++
		}
		l := len(temp)
		if lhold < l {
			lhold = l
		}
	}
	return lhold
}

 */
func lengthOfLongestSubstring(s string) int {
	var arr = make([]byte, 0)
	var i, j, res = 0, 0, 0
	for i < len(s) && j < len(s) {
		if isExist(arr, s[j]) {
			arr = arr[1:]
			i++
		} else {
			arr = append(arr, s[j])
			j++
			if res < len(arr) {
				res = len(arr)
			}
		}
	}
	return res
}
func isExist(arr []byte, key byte) bool {
	for _, value := range arr {
		if value == key {
			return true
		}
	}
	return false
}