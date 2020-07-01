package integers

import "testing"

func TestAdd(t *testing.T){

	got:=Add(2,3)
	want:=2+3
	if got!=want{
		t.Errorf("got %d want %d", got, want)
	}
}

