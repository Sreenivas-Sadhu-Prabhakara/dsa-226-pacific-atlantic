package main

import "fmt"

// PacificAtlantic returns the cells from which rain can flow to BOTH oceans
// — Pacific touching the top and left edges, Atlantic the bottom and right —
// moving only to orthogonal neighbors of equal or lower height.
// Flow is traced in reverse: climb uphill from each ocean's shoreline.
func PacificAtlantic(heights [][]int) [][2]int {
	out := [][2]int{}
	rows := len(heights)
	if rows == 0 || len(heights[0]) == 0 {
		return out
	}
	cols := len(heights[0])
	pac := make([][]bool, rows)
	atl := make([][]bool, rows)
	for r := range pac {
		pac[r] = make([]bool, cols)
		atl[r] = make([]bool, cols)
	}
	var climb func(reach [][]bool, r, c int)
	climb = func(reach [][]bool, r, c int) {
		reach[r][c] = true
		for _, d := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || reach[nr][nc] {
				continue
			}
			if heights[nr][nc] >= heights[r][c] { // reverse flow goes uphill or level
				climb(reach, nr, nc)
			}
		}
	}
	for r := 0; r < rows; r++ {
		climb(pac, r, 0)      // left edge: Pacific
		climb(atl, r, cols-1) // right edge: Atlantic
	}
	for c := 0; c < cols; c++ {
		climb(pac, 0, c)      // top edge: Pacific
		climb(atl, rows-1, c) // bottom edge: Atlantic
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pac[r][c] && atl[r][c] {
				out = append(out, [2]int{r, c})
			}
		}
	}
	return out
}

func main() {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println(PacificAtlantic(heights))
}
