package main

import "testing"

func Test0(t *testing.T) {
	if climbStairs(0) != 1 {
		t.Errorf("Test Failed")
	}
}
func Test1(t *testing.T) {
	if climbStairs(1) != 1 {
		t.Errorf("Test Failed")
	}
}

func Test2(t *testing.T) {
	if climbStairs(2) != 2 {
		t.Errorf("Test Failed")
	}
}

func Test3(t *testing.T) {
	if climbStairs(3) != 3 {
		t.Errorf("Test Failed")
	}
}
