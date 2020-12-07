package dp

func knapsack(weight []int, w int) int {
	n := len(weight)
	dp := make([][]bool, n)
	// 默认值false
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, w+1)
	}
	// 第一行的数据要特殊处理，可以利用哨兵优化
	dp[0][0] = true
	if weight[0] < w {
		dp[0][weight[0]] = true
	}

	// 动态规划
	for i := 1; i < n; i++ { // 动态规划状态转移
		for j := 0; j <= w; j++ { // 不把第i个物品放入背包
			if dp[i-1][j] {
				dp[i][j] = dp[i-1][j]
			}
		}

		for j := 0; j <= w-weight[i]; j++ { // 把第i个物品放入背包
			if dp[i-1][j] {
				dp[i][j+weight[i]] = true
			}
		}
	}

	// 最大值
	for i := w; i >= 0; i-- {
		if dp[n-1][i] {
			return i
		}
	}
	return 0
}
