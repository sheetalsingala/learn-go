package main

import "testing"

func TestRecArea(t *testing.T){
	rect:=Rectangle{10.00, 20.00}
	got:=rect.Area()
	want:=200.0
	if got != want{
		t.Errorf("Got %f want %f", got, want)
	}
}

func TestCircArea(t *testing.T){
	circ1:=Circle{10.0}
	got:=circ1.Area()
	want:=314.0
	if got != want{
		t.Errorf("Got %f want %f", got, want)
	}
}
