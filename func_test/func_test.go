package func_test

import (
	"fmt"
	"testing"
)

type TaskInfo interface {
	Process()
}

type Task struct {
	TaskId string
	X      int
	Y      int
}

func (p *Task) Process() {
	fmt.Printf("%+v\n", p)
}

func TestFuncWithStruct(t *testing.T) {
	tt := new(Task)
	tt.X = 1
	tt.Process()
	var a int
	a = 2
	fmt.Println(a)

}
