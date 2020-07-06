package main

import "testing"

func TestSearch(t * testing.T){
	dict:=Dictionary{"test":"this is a test"}
	
	t.Run("known word", func(t *testing.T){
	got, _:=dict.Search("test")
	want:="this is a test"
	
	assertStrings(t, got, want)
})

	t.Run("Unknown word", func(t *testing.T){
	_, err := dict.Search("unknown")
	want:="Could not find the word"

	if err == nil {
		t.Fatal("Expected to get an error")
	}
	
	assertStrings(t, err.Error(), want)
	})
}
func assertStrings(t *testing.T, got, want string){
	t.Helper()
	if got!= want{
		t.Errorf("got %q want %q", got, want)
	}
}
