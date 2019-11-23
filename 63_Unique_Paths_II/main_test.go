package main

import "testing"

func Test1(t *testing.T) {

	a := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	ans := uniquePathsWithObstacles(a)

	if ans != 2 {
		t.Errorf("Error")
	}
}
