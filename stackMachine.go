package main

import (
	"strings"
	"unicode"
)

type Program []Instruction

type Instruction struct {
	Opcode    Opcode
	Parameter int
}

type Opcode int

const (
	OP_PUSH Opcode = iota
	OP_ADD
	OP_JMP
	OP_RET

	// For convenience :)
	OP_IN
)

func StackMachine(p Program, x, y, z int) int {
	stack := NewStack[int]()
	pc := 0
	in := []int{x, y, z}

	for {
		switch p[pc].Opcode {
		case OP_PUSH:
			stack.Push(p[pc].Parameter)
		case OP_ADD:
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(a + b)
		case OP_JMP:
			n := stack.Pop()
			if n >= 0 {
				pc += p[pc].Parameter
			}
		case OP_RET:
			return stack.Pop()
		case OP_IN:
			stack.Push(in[p[pc].Parameter-'X'])
		}
		pc++
	}
}

func Parse(lines []string) Program {
	var p Program
	for _, l := range lines {
		var i Instruction
		fields := strings.Fields(l)
		switch fields[0] {
		case "push":
			if unicode.IsLetter(rune(fields[1][0])) {
				i.Opcode = OP_IN
				i.Parameter = int(fields[1][0])
			} else {
				i.Opcode = OP_PUSH
				i.Parameter = ParseInt(fields[1])
			}
		case "add":
			i.Opcode = OP_ADD
		case "jmpos":
			i.Opcode = OP_JMP
			i.Parameter = ParseInt(fields[1])
		case "ret":
			i.Opcode = OP_RET
		}
		p = append(p, i)
	}
	return p
}
