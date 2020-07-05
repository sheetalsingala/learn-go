package main

import "testing"

func TestArea(t *testing.T){
	checkArea:= func(t *testing.T, shape Shape, want float64){
		t.Helper()
		got:=shape.Area()
		if got != want{
			t.Errorf("Got %g want %g", got, want)
		}
	}
	
	t.Run("rectangles", func(t *testing.T){
		rect:=Rectangle{10.0, 20.0}
		checkArea(t, rect, 200.0)
	})
	
	t.Run("circles",func(t *testing.T){
		circ:=Circle{10.0}
		checkArea(t, circ, 314.0)
	})
}

