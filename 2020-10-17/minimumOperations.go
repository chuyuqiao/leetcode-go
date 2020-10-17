package _020_10_17

import "math"

func minimumOperations(leaves string) int {
	inf := math.MaxInt32
	n := len(leaves)
	f := make([][3]int, n)
	f[0][0] = bool2Int(leaves[0] == 'y')
	f[0][1] = inf
	f[0][2] = inf
	f[1][2] = inf
	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + bool2Int(leaves[0] == 'y')
		f[i][1] = min(f[i-1][0], f[i-1][1]) + bool2Int(leaves[0] == 'r')
		if i >= 2 {
			f[i][2] = min(f[i-1][1], f[i-1][2]) + bool2Int(leaves[0] == 'y')
		}
	}
	return f[n-1][2]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func bool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}
