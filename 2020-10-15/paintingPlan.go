package _020_10_15

func paintingPlan(n int, k int) int {
	res := 0
	m := map[int]int{}
	if n*n == k {
		return 1
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			tmp := i*n + j*n - i*j
			if tmp == k {
				x := combine(n, i, m)
				y := combine(n, j, m)
				res += x * y
			}
		}
	}
	return res
}

func combine(n, i int, m map[int]int) int {

	return 0
}
