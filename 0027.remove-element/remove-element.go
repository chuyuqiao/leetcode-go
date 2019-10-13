package problem0027

func removeElement(nums []int, val int) int {
	//双指针 空间复杂度O(1)
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j > 0 && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}
