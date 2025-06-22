package dp

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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

func backbag(weights []int, values []int, W int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= W; j++ {
			if weights[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}
	return dp[n][W]
}
