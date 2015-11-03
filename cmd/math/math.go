package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/honza/kids/lib"
	"strconv"
)

var args struct {
	Number int `arg:"positional,required"`
}

type Operator struct {
	name string // human readable, e.g. "+"
	f    func(a, b int) int
}

type Problem struct {
	A  int
	B  int
	Op Operator
}

func (p *Problem) Result() int {
	return p.Op.f(p.A, p.B)
}

func (p Problem) String() string {
	a := strconv.Itoa(p.A)
	b := strconv.Itoa(p.B)

	if p.A < 10 {
		a = " " + a
	}

	if p.B < 10 {
		b = " " + b
	}
	return a + " " + p.Op.name + " " + b + " ="
}

var plus = Operator{name: "+", f: func(a, b int) int {
	return a + b
}}

var minus = Operator{name: "-", f: func(a, b int) int {
	return a - b
}}

var mult = Operator{name: "*", f: func(a, b int) int {
	return a * b
}}

var Operators = []Operator{plus, minus, mult}

func RandomOperator() Operator {
	i := lib.RandomInt(3)
	return Operators[i]
}

func MakeProblem() Problem {
	op := RandomOperator()

	var a, b int

	if op.name == "+" {
		a, b = lib.RandomInt(100), lib.RandomInt(100)
	}

	if op.name == "*" {
		a, b = lib.RandomInt(10), lib.RandomInt(10)

		if a == 0 {
			a = 1
		}

		if b == 0 {
			b = 1
		}
	}

	if op.name == "-" {
		a, b = lib.RandomInt(100), lib.RandomInt(100)

		if a < b {
			a, b = b, a
		}
	}

	return Problem{A: a, B: b, Op: op}
}

func RunProblem() {
	p := MakeProblem()
	fmt.Println(p)

	for {
		n := lib.GetNumberFromUser("Answer: ")

		if n == p.Result() {
			fmt.Println("Correct!\n")
			break
		} else {
			fmt.Println("Wrong, try again\n")
		}
	}
}

func main() {
	arg.MustParse(&args)

	fmt.Printf("You have %d problems to solve.\n\n", args.Number)

	for i := 0; i < args.Number; i++ {
		RunProblem()
	}

}
