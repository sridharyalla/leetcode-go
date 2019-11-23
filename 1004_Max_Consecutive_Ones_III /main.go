package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestOnes(A []int, K int) int {
	i, j, ans, count := 0, 0, 0, 0

	for j < len(A) {
		if A[j] == 0 {
			count++
		}
		for count > K {
			if A[i] == 0 {
				count--
			}
			i++
		}
		j++
		ans = max(ans, j-i)
	}
	return ans
}

func main() {
	arr := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(longestOnes(arr, 2))
}
