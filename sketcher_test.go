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
