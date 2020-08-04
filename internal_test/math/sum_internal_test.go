package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if 3 != sum(1, 1) {
		t.Fatalf("Sum of 1 + 1 not 2")
	}
}
