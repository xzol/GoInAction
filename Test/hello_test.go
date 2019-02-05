package main

import "testing"

func TestName(t *testing.T)  {
	name := getName()
	if name != "Wor2ld!" {
		t.Error("Respone from getName is unexpected value")
	}
}
