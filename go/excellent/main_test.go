package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even5" {
		t.Errorf("expected: even, actual: %s", result)
	}
}
