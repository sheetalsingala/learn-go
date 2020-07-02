package main

import "testing"

func TestSum(t *testing.T){
	num := []int{1, 2, 4, 5}

	got:=Sum(num)
	want:=12
	
	if got!=want{
		t.Errorf("got %d want %d", got, want)
	}
}
