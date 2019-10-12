package probleam0003

func lengthOfLongestSubstring(s string) int {

	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0

	for i, v := range s {
		if location[v] >= left {
			left = location[v] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[v] = i
	}
	return maxLen
}
