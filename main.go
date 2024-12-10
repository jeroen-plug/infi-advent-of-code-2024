package main

import (
	"fmt"
	"slices"
	"sync"
)

const GridSize = 30

type Cell struct {
	X     int
	Y     int
	Z     int
	Value int
}

func (c Cell) Cmp(coords [3]int) bool {
	return c.X == coords[0] && c.Y == coords[1] && c.Z == coords[2]
}

func main() {
	p := Parse(GetLines())
	cells, sum := GetCellsAndSum(p)
	fmt.Printf("Infi part 1: %d\n", sum)
	fmt.Printf("Infi part 2: %d\n", CountClouds(cells))
}

func GetCellsAndSum(p Program) ([]Cell, int) {
	values := make(chan Cell)
	var wg sync.WaitGroup

	for x := range 30 {
		for y := range 30 {
			for z := range 30 {
				wg.Add(1)
				go func() {
					defer wg.Done()
					values <- Cell{x, y, z, StackMachine(p, x, y, z)}
				}()
			}
		}
	}

	go func() {
		wg.Wait()
		close(values)
	}()

	sum := 0
	var cells []Cell
	for v := range values {
		sum += v.Value
		if v.Value > 0 {
			cells = append(cells, v)
		}
	}
	return cells, sum
}

func CountClouds(cells []Cell) int {
	var grid [GridSize][GridSize][GridSize]int
	for _, cell := range cells {
		grid[cell.X][cell.Y][cell.Z] = cell.Value
	}

	availableCells := slices.Clone(cells)
	currentCloud := NewStack[Cell]()
	clouds := 0

	for len(availableCells) > 0 {
		currentCloud.Push(availableCells[0])
		grid[availableCells[0].X][availableCells[0].Y][availableCells[0].Z] = 0
		availableCells = removeCell(availableCells, 0)

		// Find all connected cells
		for currentCloud.Len() > 0 {
			cell := currentCloud.Pop()

			for dimension := range 3 {
				for _, delta := range []int{-1, 1} {
					coords := [3]int{cell.X, cell.Y, cell.Z}
					coords[dimension] += delta

					if coords[dimension] >= 0 && coords[dimension] < GridSize && grid[coords[0]][coords[1]][coords[2]] > 0 {
						i := slices.IndexFunc(availableCells, func(c Cell) bool { return c.Cmp(coords) })
						currentCloud.Push(availableCells[i])
						grid[coords[0]][coords[1]][coords[2]] = 0
						availableCells = removeCell(availableCells, i)
					}
				}
			}
		}
		clouds++
	}

	return clouds
}

func removeCell(cells []Cell, index int) []Cell {
	cells[index] = cells[len(cells)-1]
	return cells[:len(cells)-1]
}
