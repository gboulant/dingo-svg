package svg

import (
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

	var points []struct{ X, Y float64 } = []struct{ X, Y float64 }{
		{0.2, 0.2},
		{0.3, 0.8},
		{0.6, 0.6},
		{0.8, 0.8},
		{0.8, 0.2},
	}
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

	var points []struct{ X, Y float64 } = []struct{ X, Y float64 }{
		{0.2, 0.2},
		{0.3, 0.8},
		{0.6, 0.6},
		{0.8, 0.8},
		{0.8, 0.2},
	}
	s.Polygon(points, true)
	s.Save("output.TestSketcher_PolygonFill.svg")

	s.Clear()
	s.Polygon(points, false)
	s.Save("output.TestSketcher_PolygonVoid.svg")
}

func TestSketcher_Text(t *testing.T) {
	s := NewSketcher()

	var points []struct {
		X, Y float64
		name string
	} = []struct {
		X, Y float64
		name string
	}{
		{0.2, 0.2, "A"},
		{0.3, 0.8, "B"},
		{0.6, 0.6, "C"},
		{0.8, 0.8, "D"},
		{0.8, 0.2, "E"},
	}

	for _, p := range points {
		x, y := p.X, p.Y
		s.Point(x, y)
		s.Text(x+0.01, y+0.01, p.name)
	}

	s.Save("output.TestSketcher_Text.svg")
}

func TestSketcher_PointWithLabel(t *testing.T) {
	s := NewSketcher()

	var points []struct {
		X, Y float64
		name string
	} = []struct {
		X, Y float64
		name string
	}{
		{0.2, 0.2, "A"},
		{0.3, 0.8, "B"},
		{0.6, 0.6, "C"},
		{0.8, 0.8, "D"},
		{0.8, 0.2, "E"},
	}

	for _, p := range points {
		s.PointWithLabel(p.X, p.Y, p.name)
	}

	s.Save("output.TestSketcher_PointWithLabel.svg")
}
