package main

import (
	"strings"
	"testing"
)

func TestGetCellsAndSum(t *testing.T) {
	want := 5686200
	_, res := GetCellsAndSum(Parse(strings.Split(exampleInput, "\n")))

	if res != want {
		t.Fatalf("GetCellsAndSum() = %d, want %d", res, want)
	}
}

func TestCountCloudsExample(t *testing.T) {
	cells, _ := GetCellsAndSum(Parse(strings.Split(exampleInput, "\n")))

	want := 1
	res := CountClouds(cells)

	if res != want {
		t.Fatalf("CountClouds(example) = %d, want %d", res, want)
	}
}

func TestCountClouds(t *testing.T) {
	cells := []Cell{
		// One in each direction
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{2, 1, 1, 1},
		{1, 2, 1, 1},
		{1, 1, 2, 1},

		// Chained
		{0, 4, 0, 1},
		{0, 4, 1, 1},
		{0, 4, 2, 1},

		// Diagonal
		{1, 5, 0, 1},
	}
	want := 3
	res := CountClouds(cells)

	if res != want {
		t.Fatalf("CountClouds(simple) = %d, want %d", res, want)
	}
}

func BenchmarkGetCellsAndSum(b *testing.B) {
	p := Parse(strings.Split(exampleInput, "\n"))
	for range b.N {
		GetCellsAndSum(p)
	}
}

func BenchmarkCountClouds(b *testing.B) {
	cells, _ := GetCellsAndSum(Parse(GetLines()))
	for range b.N {
		CountClouds(cells)
	}
}
