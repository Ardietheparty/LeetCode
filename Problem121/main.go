package main

import (
	"fmt"
	"time"
)
/*
121. Best Time to Buy and Sell Stock
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
 */
func main() {
	timeStart := time.Now()
	p := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(p))

	fmt.Println(time.Since(timeStart))
}
func maxProfit(prices []int) int {
	l := len(prices)
	hol := 0
	for i := 0; i < l; i++ {
		temp := 0
		for j := i; j < l; j++ {
			if prices[j]-prices[i]>temp{
				temp = prices[j]-prices[i]
			}
		}
		if hol < temp {
			hol = temp
		}
	}
	return hol
}