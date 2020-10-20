package _020_10_20

func numberOfPatterns(m int, n int) int {
	res := 0
	used := [9]bool{}
	for k := m; k <= n; k++ {
		res += calcPatterns(-1, k, used)
		for i := 0; i < 9; i++ {
			used[i] = false
		}
	}
	return res
}

func calcPatterns(last int, k int, used [9]bool) int {
	if k == 0 {
		return 1
	}
	sum := 0
	for i := 0; i < 9; i++ {
		if isValid(last, i, used) {
			used[i] = true
			sum += calcPatterns(i, k-1, used)
			used[i] = false
		}
	}
	return sum
}

func isValid(last int, i int, used [9]bool) bool {
	if used[i] {
		return false
	}
	if last == -1 {
		return true
	}
	if (last+i)%2 == 1 {
		return true
	}
	mid := (last + i) / 2
	if mid == 4 {
		return used[mid]
	}
	if last%3 != i%3 && last/3 != i/3 {
		return true
	}
	return used[mid]
}
