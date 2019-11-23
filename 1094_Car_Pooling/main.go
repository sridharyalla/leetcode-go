package main

func carPooling(trips [][]int, capacity int) bool {
	arr := make([]int, 1001)

	for i := 0; i < len(trips); i++ {
		arr[trips[i][1]] += trips[i][0]
		arr[trips[i][2]] -= trips[i][0]
	}

	sum := 0
	for i := 0; i < len(arr); i++ {
		if sum > capacity {
			return false
		}
		sum += arr[i]
	}
	return true
}
