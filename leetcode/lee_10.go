package leetcode

func isMatch(s string, p string) bool {
	n, m := len(s) + 1, len(p) + 1
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, m)
	}
	dp[0][0] = true
	for i := 2; i < m; i += 2 {
		dp[0][i] = dp[0][i - 2] && p[i - 1] == '*'
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if p[j - 1] == '*' {
				dp[i][j] = dp[i][j - 1] || dp[i][j - 2] || (dp[i - 1][j] && s[i - 1] == p[j - 2]) || (dp[i-1][j] && p[j-2] == '.')
			} else {
				dp[i][j] = dp[i - 1][j - 1] && ((s[i - 1] == p[j - 1]) || p[j - 1] == '.')
			}
		}
	}
	return dp[n - 1][m - 1]
}