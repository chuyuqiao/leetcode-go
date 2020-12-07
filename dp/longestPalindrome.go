package dp

/**
		a  b  b  c
		0  1  2  3
    a 0 1  0  0  0
    b 1    1  1  0
    b 2       1  0
    c 3          1

  k == 0             dp[i][j] = 1
  k == 1  s[i]==s[j] dp[i][j] = 1
  other   s[i]==s[j] dp[i][j] = dp[i+1][j-1]

*/

func longestPalindrome(s string) string {
	// 初始化二维数组
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	ans := ""

	// 动态规划，状态转移
	for k := 0; k < n; k++ {
		for i := 0; i+k < n; i++ {
			j := i + k
			if k == 0 {
				dp[i][j] = 1
			} else if k == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 && k+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
