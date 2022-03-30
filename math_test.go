package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(2, 3)
	if total != 5 {
		t.Errorf("Soma(2, 3) = %d; expected 5", total)
	}
}
