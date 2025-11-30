package main

import (
	"math"
)

/*
This example is inspired from the example ch3/surface of the book "The Go
Programing Language" from Donovan and Ritchie. See source code at:

	https://github.com/adonovan/gopl.io

*/

const (
	cells = 80   // number of grid cells
	xymax = 30.0 // axis ranges (-xymax..+xymax)
)

func CardinalSine(period float64, amplitude float64) func(x, y float64) float64 {
	factor := 2 * math.Pi / period
	return func(x, y float64) float64 {
		r := factor * math.Hypot(x, y) // distance from (0,0)
		return amplitude * math.Sin(r) / r
	}
}

var f func(x, y float64) float64

func grid(i, j int) (x, y, z float64) {
	// Find point (x,y) at corner of cell (i,j).
	x = xymax * (float64(i)/cells - 0.5)
	y = xymax * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z = f(x, y)
	return x, y, z
}

func demo01() error {
	xyrange := 2 * xymax
	v := NewIsometricView(xyrange)

	period := xymax / 4.     // spatial period of the cardinal sine
	amplitude := 0.5 * xymax // amplitude of the cardianl sine
	f = CardinalSine(period, amplitude)

	for i := range cells {
		for j := range cells {
			ax, ay, az := grid(i+1, j)
			bx, by, bz := grid(i, j)
			cx, cy, cz := grid(i, j+1)
			dx, dy, dz := grid(i+1, j+1)
			v.DrawPolygon([][3]float64{
				{ax, ay, az},
				{bx, by, bz},
				{cx, cy, cz},
				{dx, dy, dz},
			}, true)
		}
	}
	return v.Save("output.demo01.svg")
}
