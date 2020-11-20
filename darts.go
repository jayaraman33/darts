package darts

import "math"

func Score(x, y float64) (score int) {
	distance := math.Sqrt(x*x + y*y)
	switch {

	case distance <= 1.0:
		score = 10
	case distance <= 5.0:
		score = 5
	case distance <= 10.0:
		score = 1
	default:
		score = 0
	}

	return score
}
