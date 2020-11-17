package sort

func bubbleSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
	return nums
}
