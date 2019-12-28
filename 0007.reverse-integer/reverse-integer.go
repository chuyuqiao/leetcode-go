package problem0007

import "math"

func reverse(x int) int {
	signed := 1
	if x < 0 {
		signed = -1
		x *= -1
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = res*10 + temp
		x = x / 10
	}

	res *= signed
	// if res < -1<<31 || res > 1<<31-1 {
	// 	res = 0
	// }
	if res < math.MinInt32 || res > math.MaxInt32 {
		res = 0
	}

	return res
}
