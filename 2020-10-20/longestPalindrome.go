package _020_10_20

func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
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
			if dp[i][j] > 0 && k+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
