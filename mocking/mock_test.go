package main

import (
	"testing"
	"bytes"
)

func TestDountdown(t *testing.T){
	buffer:= &bytes.Buffer{}
	spySleeper:=&SpySleeper{}
	Countdown(buffer, spySleeper)

	got:= buffer.String()
	want:="321Go!"

	if got!= want{
		t.Errorf("Got %q want %q", got, want)
	}
	if spySleeper.Calls!= 4{
		t.Errorf("Num sleeper calls got %d want 4", spySleeper.Calls)
	}
}
