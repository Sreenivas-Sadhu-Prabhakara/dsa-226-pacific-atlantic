package main

import (
	"reflect"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	got := PacificAtlantic(heights)
	want := [][2]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PacificAtlantic = %v, want %v", got, want)
	}

	// Single cell touches both oceans.
	if got := PacificAtlantic([][]int{{7}}); !reflect.DeepEqual(got, [][2]int{{0, 0}}) {
		t.Errorf("single cell = %v", got)
	}
	// Uniform heights: every cell reaches both.
	flat := PacificAtlantic([][]int{{1, 1}, {1, 1}})
	if len(flat) != 4 {
		t.Errorf("flat grid = %v", flat)
	}
	if got := PacificAtlantic(nil); len(got) != 0 {
		t.Errorf("empty = %v", got)
	}
}
