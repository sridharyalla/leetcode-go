package main


func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func test1 () int {
	if true {
		return 0
	} else {
		return 1
	}
}

func climbStairs(n int) int {
	dp := make([]int, Max(3, n+1))
	dp[0] = 1
		dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	
	_ = test1()
	
	return dp[n]
}
