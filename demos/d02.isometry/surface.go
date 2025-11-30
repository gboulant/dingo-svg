package main

/*
This file contains tools to create the 3d view of the surface defined by an
equation z=f(x,y). The surface is represented using an isometric view. Some
standard functions z=f(x,y) are given for examples.
*/

import "math"

// -----------------------------------------------------------
// Set of standard function builders. Call the builder with the desired
// parameters to get a Function z=f(x,y) (that can be called without parameters)

type Function func(x, y float64) float64

// CardinalSine return  the cardinal sine function with the specified spatial
// period (in the x,y direction) and z amplitude
func CardinalSine(period float64, amplitude float64) Function {
	factor := 2 * math.Pi / period
	return func(x, y float64) float64 {
		r := factor * math.Hypot(x, y) // distance from (0,0)
		return amplitude * math.Sin(r) / r
	}
}

// -----------------------------------------------------------
// Grid partition of the x,y plan

type Grid struct {
	size  int
	xymax float64
}

// Return the coordinates (x,y) of the node (i,j), considering that the xy
// values range from -xymax to xymax
func (g Grid) NodeCoordinates(i, j int) (x, y float64) {
	size := float64(g.size)
	x = g.xymax * (float64(i)/size - 0.5)
	y = g.xymax * (float64(j)/size - 0.5)
	return x, y
}

// -----------------------------------------------------------

func DrawIsometricView(f Function, gridsize int, xymax float64) *IsometricView {
	xyrange := 2 * xymax
	v := NewIsometricView(xyrange)
	g := Grid{size: gridsize, xymax: xymax}

	// xyz returns the coordinates (x,y) of the node ((i,j) on the grid and the
	// value z of the function at this grid point
	xyz := func(i, j int) (x, y, z float64) {
		x, y = g.NodeCoordinates(i, j)
		z = f(x, y)
		return x, y, z
	}

	// Create a quadrangle for each cell of the grid and draw the isometric
	// projection of this quadrangle on the canvas.
	for i := range g.size {
		for j := range g.size {
			ax, ay, az := xyz(i+1, j)
			bx, by, bz := xyz(i, j)
			cx, cy, cz := xyz(i, j+1)
			dx, dy, dz := xyz(i+1, j+1)
			v.DrawPolygon([][3]float64{
				{ax, ay, az},
				{bx, by, bz},
				{cx, cy, cz},
				{dx, dy, dz},
			}, true)
		}
	}

	return v
}
