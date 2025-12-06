package main

import (
	"math"
	"math/cmplx"

	svg "github.com/gboulant/dingo-svg"
)

// -----------------------------------------------------------
// Grid partition of the x,y plan. The grid is supposed to be mapped on a sketch
// with a centered coordinate system, whose axis range are -xymax to +xymax,
// i.e. xyrange = 2*xymax

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
// Set of holomorph complex functions

func cmplx_square(z complex128) complex128 {
	return z * z
}

func make_cmplx_rotation(angle float64) func(z complex128) complex128 {
	return func(z complex128) complex128 {
		return z * cmplx.Exp(complex(0, angle))
	}
}

func cmplx_sine(z complex128) complex128 {
	return cmplx.Sin(z)
}

func cmplx_inverse(z complex128) complex128 {
	return 1 / z
}

// make_cmplx_spiral retourne une fonction qui applique une simple rotation mais
// avec un angle d'autand plus grand que le point source est éloigné de
// l'origine. Dans la pratique, l'angle de rotation est proportionnel au module.
// On règle pour que la rotation des points éloignés de 1 soit de la valeur
// spécifiée par le paramètre angle.
func make_cmplx_spiral(angle float64) func(z complex128) complex128 {
	return func(z complex128) complex128 {
		r := cmplx.Abs(z)
		a := r * angle
		return z * cmplx.Exp(complex(0, a))
	}
}

func cmplx_func01(z complex128) complex128 {
	return z / (1 - z)
}

// -----------------------------------------------------------

// TODO: la fonction DrawFunctionGrid n'est pas très modulaire. Prévoir un
// ensemble plus modulaire de fonction, par exemple pour permettre d'afficher ou
// non la grille source (sans avoirr à ajouter des if paramétrés)

func DrawFunctionGrid(f func(z complex128) complex128, gridsize int, xymax float64) *svg.Sketcher {
	g := Grid{size: gridsize, xymax: xymax}

	// We define the two functions Zij and Zxy just to make the implementation
	// more clear/concise in the iteration loop on cells i,j.

	// Zij computes the image by f of the z value coresponding to the node i,j
	// of the grid
	Zij := func(i, j int) complex128 {
		x, y := g.NodeCoordinates(i, j)
		z := complex(x, y)
		return f(z)
	}

	// Zxy returns the coordinates x,y of a complex number, i.e. the real and
	// imaginary part respectively.
	Zxy := func(Z complex128) (x, y float64) {
		return real(Z), imag(Z)
	}

	// In this version (contrary to the "simple" implementation in the proto01
	// function), wa have to first accumulate the values of Zij on the whole
	// grid, to evaluate then the boundaries of the sketch. These boundries can
	// be different thant the xymax range that correspond to the range of the
	// source grid.

	// Initialize the Z grid
	Z := make([][]complex128, gridsize)
	for i := range Z {
		Z[i] = make([]complex128, gridsize)
	}

	// Compute the value of Z = f(z) for z values on the grid. We use these
	// iterations to evaluate the range of the image grid (min and max on xy axis)
	xmin := math.Inf(+1)
	xmax := math.Inf(-1)
	ymin := math.Inf(+1)
	ymax := math.Inf(-1)
	var x, y float64
	for i := range g.size {
		for j := range g.size {
			Z[i][j] = Zij(i, j)

			x, y = Zxy(Z[i][j])
			xmax = math.Max(x, xmax)
			xmin = math.Min(x, xmin)
			ymax = math.Max(y, ymax)
			ymin = math.Min(y, ymin)
		}
	}

	xmin = math.Min(xmin, -xymax)
	xmax = math.Max(xmax, +xymax)
	ymin = math.Min(ymin, -xymax)
	ymax = math.Max(ymax, +xymax)

	// Draw the grids (source grid and image grid)

	cnvwidth := svg.DefaultCanvasWidth
	csystem := svg.NewCoordSysWithRanges(cnvwidth, xmin, ymin, xmax, ymax)
	sk := svg.NewSketcher().WithCoordinateSystem(csystem)

	// Draw the source grid
	sk.Pencil.LineWidth = 1
	sk.Pencil.LineColor = "lightgray"
	for i := range g.size - 1 {
		for j := range g.size - 1 {
			ax, ay := g.NodeCoordinates(i+1, j)
			bx, by := g.NodeCoordinates(i, j)
			cx, cy := g.NodeCoordinates(i, j+1)
			dx, dy := g.NodeCoordinates(i+1, j+1)
			sk.Polygon([]struct{ X, Y float64 }{
				{ax, ay},
				{bx, by},
				{cx, cy},
				{dx, dy},
			}, false)
		}
	}

	// Draw the image grid
	sk.Pencil.LineWidth = 1
	sk.Pencil.LineColor = "darkgray"
	for i := range g.size - 1 {
		for j := range g.size - 1 {
			ax, ay := Zxy(Z[i+1][j])
			bx, by := Zxy(Z[i][j])
			cx, cy := Zxy(Z[i][j+1])
			dx, dy := Zxy(Z[i+1][j+1])
			sk.Polygon([]struct{ X, Y float64 }{
				{ax, ay},
				{bx, by},
				{cx, cy},
				{dx, dy},
			}, false)
		}
	}

	// We fill the cell at the top right of the grid for a better ubderstanding
	// of the transformation
	sk.Pencil.FillColor = "gray"
	i := 0
	j := 0

	ax, ay := g.NodeCoordinates(i, j)
	bx, by := g.NodeCoordinates(i+1, j)
	cx, cy := g.NodeCoordinates(i+1, j+1)
	dx, dy := g.NodeCoordinates(i, j+1)
	sk.Polygon([]struct{ X, Y float64 }{
		{ax, ay},
		{bx, by},
		{cx, cy},
		{dx, dy},
	}, true)

	ax, ay = Zxy(Z[i][j])
	bx, by = Zxy(Z[i+1][j])
	cx, cy = Zxy(Z[i+1][j+1])
	dx, dy = Zxy(Z[i][j+1])
	sk.Polygon([]struct{ X, Y float64 }{
		{ax, ay},
		{bx, by},
		{cx, cy},
		{dx, dy},
	}, true)

	return sk
}

func demo01() error {
	var f func(z complex128) complex128
	var s *svg.Sketcher

	gridsize := 20
	xymax := 4.
	f = cmplx_square
	s = DrawFunctionGrid(f, gridsize, xymax)
	s.Save("output.demo01.square.svg")

	gridsize = 20
	xymax = 2.4
	f = cmplx_sine
	s = DrawFunctionGrid(f, gridsize, xymax)
	s.Save("output.demo01.sine.svg")

	gridsize = 20
	xymax = 2.4
	f = cmplx.Cos
	s = DrawFunctionGrid(f, gridsize, xymax)
	s.Save("output.demo01.cos.svg")

	gridsize = 20
	xymax = 4.
	f = make_cmplx_rotation(math.Pi / 6)
	s = DrawFunctionGrid(f, gridsize, xymax)
	s.Save("output.demo01.rotation.svg")

	gridsize = 20
	xymax = 2.
	f = make_cmplx_spiral(math.Pi / 6)
	s = DrawFunctionGrid(f, gridsize, xymax)
	s.Save("output.demo01.spiral.svg")

	gridsize = 20
	xymax = 0.5
	f = cmplx_func01
	s = DrawFunctionGrid(f, gridsize, xymax)
	s.Save("output.demo01.func01.svg")

	return nil
}

// -----------------------------------------------------------

func main() {
	proto01()
	demo01()
}
