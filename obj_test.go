package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestObj(t *testing.T) {
	// var wg sync.WaitGroup
	var wg = sync.WaitGroup{}

	fmt.Println("Task Start")
	count := 10
	for {
		wg.Add(1)
		go func(x int, w *sync.WaitGroup) {
			fmt.Printf("input number:%v,w pointer:%p\n", x, w)
			w.Done()
		}(count, &wg)

		if count == 0 {
			break
		}
		count--
	}
	wg.Wait()
	fmt.Println("Task end")

}

type o struct {
	Num int
}

func TestValue(t *testing.T) {
	a := new(int)
	*a = 1
	fmt.Println(*a)

	// b := new(map[string]string)
	// b["test"] = "aaa"
	// fmt.Println(b)
	obj := new(o)
	obj.Num = 1
	fmt.Println(*obj)

}
