package main

import "testing"

func TestReverseInteger(t *testing.T) {
	x := 12345
	xr := reverse(x)
	if xr != 54321 {
		t.Error("expected 54321, got ", xr)
	}
	if xr == x {
		t.Error("not reversed")
	}
}
