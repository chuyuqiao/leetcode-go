package problem0007

func reverse(x int) int {
	res := 0

	for x != 0 {
		if tmp := int32(res); (tmp*10)/10 != tmp {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}

	return res
}
