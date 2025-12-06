package main

import svg "github.com/gboulant/dingo-svg"

// proto01 is a "simple" implementation (we should say straight-forward
// implementation) of a procedure to draw in the 2D complex plane the grid
// resulting from the transformation of a regular grid by application of a given
// complex function f(z).
func proto01() error {
	// Initialize a regular grid on the 2D plane
	gridsize := 20
	xymax := 2.4
	g := Grid{size: gridsize, xymax: xymax}

	// Define the function Z=f(z) to represent
	//f := cmplx_square
	//f := make_cmplx_rotation(math.Pi / 6)
	f := cmplx_sine

	// Draw the cells on a 2D sketch
	xyrange := 2 * xymax
	cnvwidth := svg.DefaultCanvasWidth
	cnvheight := svg.DefaultCanvasHeight
	csystem := svg.NewCoordSysCentered(cnvwidth, cnvheight, xyrange)
	sk := svg.NewSketcher().WithCoordinateSystem(csystem)
	sk.Pencil.LineWidth = 1
	sk.Pencil.FillColor = "whitesmoke"
	sk.Pencil.LineColor = "gray"

	// Zij computes the image of the point i,j by the function f
	Zij := func(i, j int) complex128 {
		x, y := g.NodeCoordinates(i, j)
		z := complex(x, y)
		return f(z)
	}

	// Zxy returns the coordinates x,y of a complex number, i.e. the real and
	// imaginary part respectively
	Zxy := func(Z complex128) (x, y float64) {
		return real(Z), imag(Z)
	}

	for i := range g.size {
		for j := range g.size {
			ax, ay := Zxy(Zij(i+1, j))
			bx, by := Zxy(Zij(i, j))
			cx, cy := Zxy(Zij(i, j+1))
			dx, dy := Zxy(Zij(i+1, j+1))
			sk.Polygon([]struct{ X, Y float64 }{
				{ax, ay},
				{bx, by},
				{cx, cy},
				{dx, dy},
			}, true)
		}
	}

	return sk.Save("output.proto01.svg")
}
