package main

import (
	"strings"
	"testing"
)

const exampleInput = `push 999
push X
push -3
add
jmpos 2
ret
ret
push 123
ret`

func TestStackMachine(t *testing.T) {
	want := 123
	res := StackMachine(Parse(strings.Split(exampleInput, "\n")), 7, 0, 0)

	if res != want {
		t.Fatalf("StackMachine() = %d, want %d", res, want)
	}
}
