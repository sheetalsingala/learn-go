package main

import "testing"


func TestHello2(t *testing.T){

	assertMessage:=func(t *testing.T, got, want string){
		t.Helper()
		if got!=want{
			t.Errorf("got %q want %q", got, want)
		}
	}
	
	t.Run("People", func(t *testing.T){
		got:=Hello2("Anne")
		want:="Hello, Anne"
		assertMessage(t, got, want)
	})
	
	t.Run("World", func(t *testing.T){
		got:=Hello2("")
		want:="Hello, World"
		assertMessage(t, got, want)

	})
}


