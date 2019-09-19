package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	arr := [...]int{1, 2, 3}

	for _, e := range arr {
		t.Log(e)
	}

	t.Log(arr[:len(arr)])
}

func TestArrayVsSlice(t *testing.T) {
	var x [3]int = [3]int{1, 2, 3}
	var y []int = []int{4, 5, 6}
	t.Log(x, y)
	z := x
	z[0] = 999
	d := y
	d[0] = 998

	t.Log(x, y)
}
