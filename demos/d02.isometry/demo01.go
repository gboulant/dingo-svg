package main

import (
	"math"

	svg "github.com/gboulant/dingo-svg"
)

/*
This example is inspired from the example ch3/surface of the book "The Go
Programing Language" from Donovan and Ritchie. See source code at:

	https://github.com/adonovan/gopl.io

*/

const (
	cells   = 100 // number of grid cells
	xyrange = 1.0 // axis ranges (-xyrange..+xyrange)
)

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := 0.5 + (x-y)*cos30
	sy := 0.5 + z*0.4 - (x+y)*sin30
	return sx, sy
}

func f(x, y float64) float64 {
	a := 30.
	r := math.Hypot(a*x, a*y) // distance from (0,0)
	return math.Sin(r) / r
}

func demo01() error {
	s := svg.NewSketcher()
	s.Pencil.LineWidth = 1
	s.Pencil.FillColor = "whitesmoke"
	s.Pencil.LineColor = "gray"

	for i := range cells {
		for j := range cells {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			s.Quadrangle(ax, ay, bx, by, cx, cy, dx, dy, true)
		}
	}
	return s.Save("output.demo01.svg")
}
