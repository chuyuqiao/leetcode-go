package _020_10_20

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
