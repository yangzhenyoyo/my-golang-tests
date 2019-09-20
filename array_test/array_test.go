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

func TestSliceCopy(t *testing.T) {
	var s []string = []string{"a", "b", "c", "d", "e"}
	t.Log("initial slice:", s)
	ss := s[:3]
	t.Log("Fetch 3 elements:", ss)
	ss[2] = "cc"
	t.Log("set index:2 in ss,then the s:", s)

	sss := make([]string, len(ss))
	ncp := copy(sss, ss)
	t.Log("copy nums:", ncp)
	t.Log("copy ss to sss:", sss)
	sss[2] = "ccc"
	t.Logf("s:%v,ss:%v,sss:%v\n", s, ss, sss)
}

func TestSliceAsFuncParam(t *testing.T) {
	// slice的传参，值传递依然可以修改值,注意坑
	myslice := []int{1, 2, 3, 4}

	func(s []int) {
		if s == nil {
			panic("slice need initial")
		}
		s[2] = 333
	}(myslice)

	t.Log("slice result:", myslice)
}

func TestArrAsFuncParam(t *testing.T) {
	//数组需要引用传递才可能让函数修改值
	myarr := [3]int{1, 2, 3}

	func(s [3]int) {
		s[0] = 11
	}(myarr)
	t.Log(myarr)
	func(s *[3]int) {
		s[0] = 111
	}(&myarr)

	t.Log(myarr)
}
