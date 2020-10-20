package _020_10_20

func permute(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	out := make([]int, n)
	for i, v := range nums {
		out[i] = v
	}
	backtrack(0, n, out, &res)
	return res
}

func backtrack(first int, n int, out []int, res *[][]int) {
	if first == n {
		*res = append(*res, append([]int{}, out...))
	}
	for i := first; i < n; i++ {
		out[first], out[i] = out[i], out[first]
		backtrack(first+1, n, out, res)
		out[first], out[i] = out[i], out[first]
	}
}
