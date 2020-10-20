package _020_10_20

func maxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1
	ans := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
