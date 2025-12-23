package wfc

import "errors"

// Possible adjacent tile indicies for each tile index.
type AdjacencyRules map[int]struct{
	Up		[]int
	Right	[]int
	Down	[]int
	Left	[]int
}

// Possibility (weight) of each tile index.
type FrequencyHints map[int]int

type Grid2DSize struct {
	Width 	int
	Height	int
}

/*
Populates grid with tile indicies in-place according to the rules present in the input arguments. 
Returns error if something if something went wrong.

If grid is of size 0 (in width or height), no error is returned (nil).
*/
func WfcGenerate(
	grid				*[][]int,
	adjacencyRules		AdjacencyRules,
	frequencyHints		FrequencyHints,
	gridSize			Grid2DSize,
) error {
	if len(adjacencyRules) != len(frequencyHints) {
		return errors.New("adjacencyRules length doesn't match frequencyHints length")
	}
	gw := gridSize.Width
	gh := gridSize.Height

	// number of possible indicies
	ni := len(adjacencyRules) 	

	indicies := make([]int, ni)
	for k, _ := range adjacencyRules {
		indicies = append(indicies, k)
	}

	// pre-populate entire wave
	wave := make([][][]int, gh)
	for row := range wave {
		wave[row] = make([][]int, gw)
		for cell := range wave[row] {
			wave[row][cell] = indicies
		}
	}
	gSize := gridSize.Width * gridSize.Height
	minEntropyCells := []int{}

	return nil
}
