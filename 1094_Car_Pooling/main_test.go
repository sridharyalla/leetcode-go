package main

import "testing"

func Test1(t *testing.T) {
	trips := [][]int{
		{2, 1, 5},
		{3, 3, 7},
	}
	ans := carPooling(trips, 4)

	if ans != false {
		t.Errorf("Error")
	}
}
