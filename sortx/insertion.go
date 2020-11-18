package sortx

import "fmt"

func insertionSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
		fmt.Printf("%d=>%v \n", i, nums)
	}
	return nums
}
