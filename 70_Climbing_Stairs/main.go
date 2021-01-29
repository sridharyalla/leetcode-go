package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func climbStairs(n int) int {
	dp := make([]int, max(3, n+1))
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs(n int) int {
	dp := make([]int, max(3, n+1))
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
