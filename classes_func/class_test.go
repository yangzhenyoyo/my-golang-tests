package classesfunc

import (
	"fmt"
	"testing"
)

type Class interface {
	Do() string
}

var (
	f = make(map[string]func() Class)
)

func Set(name string, factory func() Class) {
	f[name] = factory
}

func Get(name string) Class {
	if ff, ok := f[name]; ok {
		return ff()
	}
	panic("nanme not found")

}

type Class1 struct {
}

func (cls1 *Class1) Do() string {
	return fmt.Sprintln("this class1")
}

type Class2 struct {
}

func (cls2 *Class2) Do() string {
	return fmt.Sprintln("this class2")
}

func TestFactoryClass(t *testing.T) {
	Set("Class1", func() Class { return new(Class1) })
	Set("Class2", func() Class { return new(Class2) })

	var cls1, cls2 Class
	cls1 = Get("Class1")

	fmt.Printf("%T ,%v\n", cls1, cls1.Do())
	cls2 = Get("Class2")

	fmt.Printf("%T ,%v\n", cls2, cls2.Do())
}
