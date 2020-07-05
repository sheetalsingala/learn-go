package main

import "testing"

func TestPerimeter(t *testing.T){
	rect:= Rectangle{10.00, 20.00}	
	got:= Perimeter(rect)
	want:=60.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T){
	rect:=Rectangle{10.00, 20.00}
	got:=Area(rect)
	want:=200.0
	if got != want{
		t.Errorf("got %f want %f", got, want)
	}
}
