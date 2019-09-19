package slice_test

import "testing"

func TestLoop(t *testing.T) {
	n := 0

	for {
		if n > 10 {
			break
		}
		t.Log("this is", n)
		n++

	}
}
