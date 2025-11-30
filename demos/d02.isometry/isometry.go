package main

import (
	"math"

	svg "github.com/gboulant/dingo-svg"
)

const angle = math.Pi / 6 // angle of x, y axes (=30°)
const zscale = 0.82       // sqrt(2/3)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func isometric(x, y, z float64) (X, Y float64) {
	X = (y - x) * cos30
	Y = z*zscale - (x+y)*sin30
	return X, Y
}

type IsometricView struct {
	sk *svg.Sketcher
}

// xyrange is the axis ranges (-xyrange/2, +xyrange/2)
func NewIsometricView(xyrange float64) *IsometricView {
	cnvwidth := svg.DefaultCanvasWidth
	cnvheight := svg.DefaultCanvasHeight
	csystem := svg.NewCoordSysCentered(cnvwidth, cnvheight, xyrange)
	sk := svg.NewSketcher().WithCoordinateSystem(csystem)
	sk.Pencil.LineWidth = 1
	sk.Pencil.FillColor = "whitesmoke"
	sk.Pencil.LineColor = "gray"
	return &IsometricView{sk}
}

func (v IsometricView) DrawLine(A, B [3]float64) {
	x, y, z := A[0], A[1], A[2]
	XA, YA := isometric(x, y, z)
	x, y, z = B[0], B[1], B[2]
	XB, YB := isometric(x, y, z)
	v.sk.Edge(XA, YA, XB, YB)
}

func (v IsometricView) DrawAxis() {
	O := [3]float64{0., 0., 0.}
	I := [3]float64{1., 0., 0.}
	J := [3]float64{0., 1., 0.}
	K := [3]float64{0., 0., 1.}
	color := v.sk.Pencil.LineColor
	v.sk.Pencil.LineColor = "red"
	v.DrawLine(O, I)
	v.sk.Pencil.LineColor = "green"
	v.DrawLine(O, J)
	v.sk.Pencil.LineColor = "blue"
	v.DrawLine(O, K)
	v.sk.Pencil.LineColor = color
}

func (v IsometricView) DrawPolygon(polygon3d [][3]float64, fill bool) {
	polygon2d := make([]struct{ X, Y float64 }, len(polygon3d))
	for i, p := range polygon3d {
		x, y, z := p[0], p[1], p[2]
		X, Y := isometric(x, y, z)
		polygon2d[i] = struct{ X, Y float64 }{X, Y}
	}

	v.sk.Polygon(polygon2d, fill)
}

func (v IsometricView) Save(svgpath string) error {
	return v.sk.Save(svgpath)
}
