package sort

func insertionSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	for i := 1; i < n; i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}
	return nums
}
