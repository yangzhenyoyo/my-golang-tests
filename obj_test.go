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
	// *a = 1
	fmt.Println(*a)

	// b := new(map[string]string)
	// *b["test"] = "aaa"
	// fmt.Println(b)
	// obj := new(o)
	// obj.Num = 1
	// fmt.Println(*obj)

}

func TestMapValue(t *testing.T){

	m :=make(map[string]string)

	m["name"]="jj"
	func(m2 map[string]string){
     m2["name"]="dd"
	}(m)

	fmt.Println(m)

}

func TestChanValue(t *testing.T){
	ch := make(chan string,100)

	func(c chan string){
      ch<-"hello"
	}(ch)
  fmt.Println(<-ch)
}



func TestStruct(t *testing.T){

	 s := o{Num:1}
	 
	 func(st *o){
      st.Num=2
	 }(&s)
	 fmt.Println(s)
 
}