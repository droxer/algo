package dp

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 1 {
		return 2
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
