package _020_10_12

import "sort"

func breakfastNumber(staple []int, drinks []int, x int) int {
	const NUMBER = 1000000007
	count := 0
	sort.Ints(staple)
	sort.Ints(drinks)
	for i := 0; i < len(staple); i++ {
		if staple[i] > x {
			break
		}
		low, high := 0, len(drinks)
		target := x - staple[i]
		for j := len(drinks) - 1; j >= 0 && low < high; j-- {
			mid := (low + high) / 2
			if drinks[mid] > target {
				high = mid
			} else {
				low = mid + 1
			}
		}
		count += low
	}
	return count % NUMBER
}
