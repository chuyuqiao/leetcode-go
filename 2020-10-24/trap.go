package _020_10_24

func trap(height []int) int {
	n, ans := len(height), 0
	for i := 0; i < n; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		for j := i; j < n; j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		ans += min(maxLeft, maxRight) - height[i]
	}
	return ans
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
