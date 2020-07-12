package problem0006

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 1 {
		return s
	}
	sl := make([]string, numRows)
	dir := 1
	curr := 0
	for _, v := range s {
		sl[curr] = sl[curr] + string(v)
		curr += dir
		if curr == 0 || curr == numRows-1 {
			dir = -dir
		}
	}
	slr := ""
	for _, v := range sl {
		slr += v
	}
	return slr
}
