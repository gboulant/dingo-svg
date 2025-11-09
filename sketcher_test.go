package svg

import (
	"fmt"
	"testing"
)

const output_TestSketcher_LineTo string = `<svg xmlns='http://www.w3.org/2000/svg' width='600' height='600'>
<line x1='480.00' y1='120.00' x2='480.00' y2='480.00' style='stroke: black; stroke-width: 2; fill: black'/>
<line x1='480.00' y1='480.00' x2='120.00' y2='480.00' style='stroke: black; stroke-width: 2; fill: black'/>
<line x1='120.00' y1='480.00' x2='120.00' y2='120.00' style='stroke: black; stroke-width: 2; fill: black'/>
<line x1='120.00' y1='120.00' x2='480.00' y2='120.00' style='stroke: black; stroke-width: 2; fill: black'/>
</svg>`

func TestSketcher_LineTo(t *testing.T) {
	s := NewSketcher()
	s.MoveTo(0.8, 0.8)
	s.LineTo(0.8, 0.2)
	s.LineTo(0.2, 0.2)
	s.LineTo(0.2, 0.8)
	s.LineTo(0.8, 0.8)
	s.Save("output.TestSketcher_LineTo.svg")
	res := s.ToSVG()
	ref := output_TestSketcher_LineTo
	if res != ref {
		t.Errorf("result is:\n%s\nShould be:\n%s", res, ref)
	}
}

const output_TestSketcher_Circle string = `<svg xmlns='http://www.w3.org/2000/svg' width='600' height='600'>
<line x1='120.00' y1='480.00' x2='480.00' y2='120.00' style='stroke: black; stroke-width: 2; fill: black'/>
<circle cx='480.00' cy='120.00' r='60.00' style='stroke: black; stroke-width: 2; fill: black'/>
<line x1='480.00' y1='120.00' x2='480.00' y2='480.00' style='stroke: black; stroke-width: 2; fill: black'/>
<circle cx='480.00' cy='480.00' r='60.00' style='stroke: black; stroke-width: 2; fill: none'/>
</svg>`

func TestSketcher_Circle(t *testing.T) {
	s := NewSketcher()
	s.MoveTo(0.2, 0.2)
	s.LineTo(0.8, 0.8)
	cx, cy := s.Position()
	s.Circle(cx, cy, 0.1, true)
	s.LineTo(0.8, 0.2)
	cx, cy = s.Position()
	s.Circle(cx, cy, 0.1, false)
	s.Save("output.TestSketcher_Circle.svg")

	res := s.ToSVG()
	ref := output_TestSketcher_Circle
	if res != ref {
		t.Errorf("result is:\n%s\nShould be:\n%s", res, ref)
	}
}

func TestSketcher_Point(t *testing.T) {
	s := NewSketcher()
	s.Point(0.2, 0.2)
	s.LineTo(0.8, 0.8)
	x, y := s.Position()
	s.Circle(x, y, 0.1, false)
	s.Save("output.TestSketcher_Point.svg")
}

func TestSketcher_Polyline(t *testing.T) {
	s := NewSketcher()

	points := testpoints()
	s.Polyline(points, true)
	s.Save("output.TestSketcher_Polyline.svg")
}

func TestSketcher_BaseShapes(t *testing.T) {
	s := NewSketcher()

	x0, y0 := 0.1, 0.2
	w, h := 0.4, 0.5

	x, y := x0, y0
	s.Triangle(x, y, x+w*0.5, y+h, x+w, y, false)
	s.Text(x, y, "Triangle Vide")

	x, y = x0+0.1, y0+0.1
	s.Triangle(x, y, x+w*0.5, y+h, x+w, y, true)
	s.Pencil.FontColor = "white"
	s.Text(x, y, "Triangle Plein")

	s.Pencil.LineColor = "red"
	s.Pencil.FillColor = "blue"
	x, y = x0+0.4, y0-0.1
	s.Quadrangle(x, y, x+w*0.6, y+h*0.2, x+w*0.7, y+h*0.8, x, y+h, true)
	s.Pencil.FontColor = s.Pencil.LineColor
	s.Text(x, y, "Quadrangle Plein")

	s.Pencil.LineColor = "green"
	x, y = x0+0.5, y0
	s.Quadrangle(x, y, x+w*0.6, y+h*0.2, x+w*0.7, y+h*0.8, x, y+h, false)
	s.Pencil.FontColor = s.Pencil.LineColor
	s.Text(x, y, "Quadrangle Vide")

	s.Pencil.LineColor = "yellow"
	x, y = x0+0.3, y0+0.7
	s.Rectangle(x, y, w, h, false)
	s.Pencil.FontColor = s.Pencil.LineColor
	s.Text(x, y, "Rectangle Vide")

	s.Save("output.TestSketcher_BaseShapes.svg")
}

func TestSketcher_Polygon(t *testing.T) {
	s := NewSketcher()

	points := testpoints()
	s.Polygon(points, true)
	s.Save("output.TestSketcher_PolygonFill.svg")

	s.Clear()
	s.Polygon(points, false)
	s.Save("output.TestSketcher_PolygonVoid.svg")
}

func TestSketcher_Text(t *testing.T) {
	s := NewSketcher()

	points := testpointsWithNames()
	for _, p := range points {
		x, y := p.X, p.Y
		s.Point(x, y)
		s.Text(x+0.01, y+0.01, p.name)
	}

	s.Save("output.TestSketcher_Text.svg")
}

func TestSketcher_PointWithLabel(t *testing.T) {
	s := NewSketcher()

	points := testpointsWithNames()
	for _, p := range points {
		s.PointWithLabel(p.X, p.Y, p.name)
	}

	s.Save("output.TestSketcher_PointWithLabel.svg")
}

func TestSketcher_BottomLeftCoordinateSystem(t *testing.T) {
	cnvwidth := DefaultCanvasWidth
	cnvheight := DefaultCanvasHeight
	xrange := 10.
	cs := NewCoordSysBottomLeft(cnvwidth, cnvheight, xrange)
	s := NewSketcher().WithCoordinateSystem(cs)

	points := testpoints()
	s.Polygon(points, true)

	scaling := xrange
	for i, p := range points {
		points[i].X = scaling * p.X
		points[i].Y = scaling * p.Y
	}

	s.Polygon(points, false)

	s.Save("output.TestSketcher_BottomLeftCoordinateSystem.svg")
}

func TestSketcher_CenteredCoordinateSystem(t *testing.T) {
	cnvwidth := DefaultCanvasWidth
	cnvheight := DefaultCanvasHeight
	xrange := 2. // define an xrange of 2 to range from -width to width
	// the xrange is the range from the left boudary of the canvas to the right boundary
	cs := NewCoordSysCentered(cnvwidth, cnvheight, xrange)
	s := NewSketcher().WithCoordinateSystem(cs)

	points := testpoints()
	s.Polygon(points, true)

	// Center symetry by Origin
	for i, p := range points {
		points[i].X = -1 * p.X
		points[i].Y = -1 * p.Y
	}

	s.Polygon(points, false)

	s.Save("output.TestSketcher_CenteredCoordinateSystem.svg")
}

func TestSketcher_BoundedByCoordinateSystem(t *testing.T) {
	cnvwidth := DefaultCanvasWidth
	points := testpoints()
	// translate points
	tx := 10.
	ty := -5.
	for i, p := range points {
		points[i].X = p.X + tx
		points[i].Y = p.Y + ty
	}

	xoffset := 0.05
	yoffset := 0.05
	cs := NewCoordSysBoundedBy(cnvwidth, points, xoffset, yoffset)
	s := NewSketcher().WithCoordinateSystem(cs)

	s.Polygon(points, false)

	var label string
	for i, p := range points {
		label = fmt.Sprintf("p%d (%.2f, %.2f)", i, p.X, p.Y)
		s.PointWithLabel(p.X, p.Y, label)
	}

	s.Save("output.TestSketcher_BoundedByCoordinateSystem.svg")
}
