package main

import "math/rand"

func dataset_random(xmin, xmax, ymin, ymax float64, size int) []Vector {
	vectors := make([]Vector, size)
	xrange := xmax - xmin
	yrange := ymax - ymin
	for i := range size {
		xi := xmin + xrange*rand.Float64()
		yi := ymin + yrange*rand.Float64()
		vectors[i] = Vector{X: xi, Y: yi}
	}
	return vectors
}
