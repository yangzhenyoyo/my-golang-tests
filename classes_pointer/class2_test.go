package classespointer

import (
	"fmt"
	"testing"
)

type Class interface {
	Done() string
}

var f = make(map[string]Class)

func Set(name string, p Class) {
	f[name] = p
}

func Get(name string) Class {
	if p, ok := f[name]; ok {
		return p
	}
	panic("name not found")
}

//Class1
type Class1 struct{}

func (cls *Class1) Done() string {
	return fmt.Sprintln("this is Class1")
}

//Class2
type Class2 struct{}

func (cls *Class2) Done() string {
	return fmt.Sprintln("this is Class2")
}

func parser(p interface{}) (interface{}, string) {
	var remark string

	switch p.(type) {
	case int:
		remark = "type int"
	case string:
		remark = "type string"
	default:
		remark = "unknown type"
	}
	return p, remark
}

func TestFactory2(t *testing.T) {
	Set("Class1", new(Class1))
	Set("Class2", new(Class2))
	cls1 := Get("Class1")
	cls2 := Get("Class2")
	fmt.Printf("%T,%v\n", cls1, cls1.Done())
	fmt.Printf("%T,%v\n", cls2, cls2.Done())
	val := "sb"
	a, b := parser(val)
	fmt.Printf("parse %v return result %v", b, a)
}
