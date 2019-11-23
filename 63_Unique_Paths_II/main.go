package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 0 {
		obstacleGrid[0][0] = 1
	} else {
		obstacleGrid[0][0] = 0
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[0][i] == 0 {
			obstacleGrid[0][i] = obstacleGrid[0][i-1]
		} else {
			obstacleGrid[0][i] = 0
		}
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[i][0] == 0 {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		} else {
			obstacleGrid[i][0] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}

	return obstacleGrid[n-1][m-1]
}
