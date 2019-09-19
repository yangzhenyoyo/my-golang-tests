package interfaces

import (
	"fmt"
	"testing"
)

type Code string

type Program interface {
	Run() Code
}

type JavaProgram struct {
}

func (j *JavaProgram) Run() Code {
	return "Run Java"
}

type GoProgram struct {
}

func (g *GoProgram) Run() Code {
	return "Run Go"
}

func RunCode(p Program) {
	fmt.Printf("%T %v\n", p, p.Run())
}

func TestRunProgram(t *testing.T) {
	goPro := new(GoProgram)
	javaPro := new(JavaProgram)
	RunCode(goPro)
	RunCode(javaPro)
}
