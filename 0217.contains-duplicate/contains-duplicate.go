package problem0217

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for i, n := range nums {
		if _, prs := m[n]; prs {
			return true
		}
		m[n] = i
	}

	return false
}
