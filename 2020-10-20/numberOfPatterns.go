package _020_10_20

func numberOfPatterns(m int, n int) int {
	used := [9]bool{}
	res := 0
	for k := m; k <= n; k++ {
		res += calcPatterns(-1, k, used)
		for i := 0; i < 9; i++ {
			used[i] = false
		}
	}
	return res
}

func calcPatterns(last int, k int, used [9]bool) int {
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
	//1.被使用
	if used[i] {
		return false
	}
	//2.首次
	if last == -1 {
		return true
	}
	//3.同行同列相邻，象棋马
	if (last+i)%2 == 1 {
		return true
	}
	//4.中间
	mid := (last + i) / 2
	if mid == 4 {
		return used[mid]
	}
	//5.对角线
	if last/3 != i/3 && last%3 != i%3 {
		return true
	}
	//6.同行同列不相邻
	return used[mid]
}
