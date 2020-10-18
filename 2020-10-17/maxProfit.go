package _020_10_17

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			if prices[i]-minPrice > maxProfit {
				maxProfit = prices[i] - minPrice
			}
		}
	}
	return maxProfit
}
