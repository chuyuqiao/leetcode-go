package sortx

import "fmt"

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		flag := false
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
		fmt.Printf("%d=>%v \n", i, nums)
	}
	return nums
}
