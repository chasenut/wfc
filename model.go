package wfc

import "math"

type Iterator interface {
	Iterate(n int) ([][]int, bool, bool)
}

type Generator interface {
	Generate() ([][]int, bool)
}

type ExecBlueprint interface {
	Iterator
	Generator
	InLimits(x, y int) bool
	Propagate() bool
	Clear()
}

type BaseModel struct {
	Wave 					[][][]bool		// Superpositions of each cell (x, y) - all patterns that could be in fit in such cell
	Changes 				[][]bool		// 2D array of collapsed cells
	Weights					[]float64		// Frequency hints for each tile index
	Periodic				bool			// If the input image should be periodic (tesselated surface)
	GenerationSuccessful	bool			// Whether generation was successful
	T						int				// Count of possible patterns
	Fx, Fy					int				// Final width and height of output
	RngSet					bool			// Random number generator set by user
	Rng						func() float64	// Random number generator used during generation
}

// Returns true if generation is successful and complete, false otherwise
func (baseModel *BaseModel) Observe(specificModel ExecBlueprint) bool {
	min := 1000.0
	cellminx := -1
	cellminy := -1
	distribution := make([]float64, baseModel.T)
	
	for x := 0; x < baseModel.Fx; x++ {
		for y := 0; y < baseModel.Fy; y++ {
			if !specificModel.InLimits(x, y) {
				continue
			}

			sum := 0.0
			for t := 0; t < baseModel.T; t++ {
				if baseModel.Wave[x][y][t] {
					distribution[t] = baseModel.Weights[t]
				} else {
					distribution[t] = 0.0
				}
				sum += distribution[t]
			}

			if sum == 0.0 {
				baseModel.GenerationSuccessful = false
				return true
			}

			for t := 0; t < baseModel.T; t++ {
				distribution[t] /= sum
			}

			entropy := 0.0

			for i := 0; i < len(distribution); i++ {
				if distribution[i] > 0.0 {
					entropy += -distribution[i] * math.Log(distribution[i])
				}
			}

			noise := 0.000001 * baseModel.Rng()

			if entropy > 0 && entropy + noise < min {
				min = entropy + noise
				cellminx = x
				cellminy = y
			}
		}
	}

	if cellminx == -1 && cellminy == -1 {
		baseModel.GenerationSuccessful = true
		return true // finished with success
	}

	for t := 0; t < baseModel.T; t++ {
		if baseModel.Wave[cellminx][cellminy][t] {
			distribution[t] = baseModel.Weights[t]
		} else {
			distribution[t] = 0.0
		}
	}

	r := randomIndex(distribution, baseModel.Rng())

	for t := 0; t < baseModel.T; t++ {
		baseModel.Wave[cellminx][cellminy][t] = (r == t)
	}

	baseModel.Changes[cellminx][cellminy] = true

	return false // not finished yet
}
