package strx

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	rk, ans := -1, 0
	n := len(s)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]] = 1
			rk++
		}
		ans = max(rk-i+1, ans)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
