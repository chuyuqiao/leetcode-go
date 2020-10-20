package _020_10_20

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			if price-minPrice > maxProfit {
				maxProfit = price - minPrice
			}
		}
	}
	return maxProfit
}
