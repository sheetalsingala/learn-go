// Test file is of type xxx_test.go 
// Test function must begin with TestXxx and takes one argument t *testing.T

package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Chris")
	want := "Hello, Chris"
	
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}
