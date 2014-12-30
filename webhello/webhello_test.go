package main

import "testing"

func TestIncrementAndGet(t *testing.T) {

	n := incrementAndGet()

	if n != 1 {
		t.Fail()
	}

	n = incrementAndGet()
	if n != 2 {
		t.Fail()
	}

	n = incrementAndGet()
	if n != 3 {
		t.Fail()
	}
}
