package svg

import (
	"fmt"
	"os"
)

const (
	headPattern  = "<svg xmlns='http://www.w3.org/2000/svg' width='%d' height='%d'>"
	linePattern  = "<line x1='%.2f' y1='%.2f' x2='%.2f' y2='%.2f' style='%s'/>"
	textPattern  = "<text x='%.2f' y='%.2f' style='%s'>%s</text>"
	rectPattern  = "<rect x='%.2f' y='%.2f' width='%.2f' height='%.2f' style='%s'/>"
	circPattern  = "<circle cx='%.2f' cy='%.2f' r='%.2f' style='%s'/>"
	polygPattern = "<polygon points='%s' style='%s'/>"
	footPattern  = "</svg>"
)

// ===========================================================================
// Sketch builder
// ===========================================================================

var defaultCoordinateSystem = NewCoordinateSystem()

type Sketcher struct {
	x, y   float64
	body   string
	cs     *CoordinateSystem
	Pencil *Pencil
}

func NewSketcher() *Sketcher {
	cs := defaultCoordinateSystem
	return &Sketcher{x: 0., y: 0, body: "", cs: cs, Pencil: defaultPencil.Clone()}
}

func (s *Sketcher) WithCoordinateSystem(cs *CoordinateSystem) *Sketcher {
	s.cs = cs
	return s
}

// --------------------------------------------------------------------
// Sketch export and display functions

func (s Sketcher) ToSVG() string {
	svg := fmt.Sprintf(headPattern, 600, 600) + "\n"
	svg += s.body
	svg += footPattern
	return svg
}

func (s Sketcher) String() string {
	return s.ToSVG()
}

func (s Sketcher) Save(svgpath string) error {
	file, err := os.OpenFile(svgpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(s.ToSVG())
	if err != nil {
		return err
	}
	return nil
}

// --------------------------------------------------------------------
// Sketch management functions

func (s *Sketcher) Clear() {
	s.body = ""
}

func (s Sketcher) Position() (x, y float64) {
	return s.x, s.y
}

func (s Sketcher) canvasCoordinates(x, y float64) (px, py float64) {
	return s.cs.canvasCoordinates(x, y)
}

func (s Sketcher) canvasScaling(size float64) (psize float64) {
	return s.cs.canvasScaling(size)
}

const factor = 1.

func (s Sketcher) pointSize() float64 {
	return factor * float64(s.Pencil.LineWidth) / float64(s.cs.cnvxsize)
}

// --------------------------------------------------------------------
// Turtle-like drawing functions

func (s *Sketcher) MoveTo(x, y float64) {
	s.x = x
	s.y = y
}

func (s *Sketcher) LineTo(x, y float64) {
	px1, py1 := s.canvasCoordinates(s.x, s.y)
	px2, py2 := s.canvasCoordinates(x, y)
	style := s.Pencil.DrawStyle()
	s.body += fmt.Sprintf(linePattern, px1, py1, px2, py2, style) + "\n"
	s.x = x
	s.y = y
}

// --------------------------------------------------------------------
// Function to create primitive shapes

func (s *Sketcher) Point(x, y float64) {
	r := s.pointSize()
	s.Circle(x, y, r, true)
}

func (s *Sketcher) Edge(x1, y1, x2, y2 float64) {
	s.MoveTo(x1, y1)
	s.LineTo(x2, y2)
}

func (s *Sketcher) Circle(cx, cy, r float64, fill bool) {
	pcx, pcy := s.canvasCoordinates(cx, cy)
	pr := s.canvasScaling(r)
	style := s.Pencil.DrawStyleWithFillMode(fill)
	s.body += fmt.Sprintf(circPattern, pcx, pcy, pr, style) + "\n"
	s.x = cx
	s.y = cy
}

func (s *Sketcher) Triangle(x1, y1, x2, y2, x3, y3 float64, fill bool) {
	px1, py1 := s.canvasCoordinates(x1, y1)
	px2, py2 := s.canvasCoordinates(x2, y2)
	px3, py3 := s.canvasCoordinates(x3, y3)
	coords := fmt.Sprintf("%.2f,%.2f %.2f,%.2f %.2f,%.2f", px1, py1, px2, py2, px3, py3)
	style := s.Pencil.DrawStyleWithFillMode(fill)
	s.body += fmt.Sprintf(polygPattern, coords, style) + "\n"
	s.x = x3
	s.y = y3
}

func (s *Sketcher) Quadrangle(x1, y1, x2, y2, x3, y3, x4, y4 float64, fill bool) {
	px1, py1 := s.canvasCoordinates(x1, y1)
	px2, py2 := s.canvasCoordinates(x2, y2)
	px3, py3 := s.canvasCoordinates(x3, y3)
	px4, py4 := s.canvasCoordinates(x4, y4)
	coords := fmt.Sprintf("%.2f,%.2f %.2f,%.2f %.2f,%.2f %.2f,%.2f", px1, py1, px2, py2, px3, py3, px4, py4)
	style := s.Pencil.DrawStyleWithFillMode(fill)
	s.body += fmt.Sprintf(polygPattern, coords, style) + "\n" // this concatenation is time consuming
	s.x = x4
	s.y = y4
}

func (s *Sketcher) Rectangle(x, y, width, height float64, fill bool) {
	px, py := s.canvasCoordinates(x, y)
	pw := s.canvasScaling(width)
	ph := s.canvasScaling(height)
	style := s.Pencil.DrawStyleWithFillMode(fill)
	s.body += fmt.Sprintf(rectPattern, px, py, pw, ph, style) + "\n" // this concatenation is time consuming
	s.x = x
	s.y = y
}

// Polygon draw a polygon, i.e. closed polyline defined by an ordered
// set of points related by edges. The variable points is a list of
// point coordinates, each point coordinates is a tuple (x,y).
func (s *Sketcher) Polygon(points []struct{ X, Y float64 }, fill bool) {
	var x, y float64
	var coords string
	for _, p := range points {
		x, y = p.X, p.Y
		px, py := s.canvasCoordinates(x, y)
		coords += fmt.Sprintf("%.2f,%.2f ", px, py)
	}
	style := s.Pencil.DrawStyleWithFillMode(fill)
	s.body += fmt.Sprintf(polygPattern, coords, style) + "\n"
	s.x = x
	s.y = y
}

// Polyline draws a continuous line made of multiple connected edges,
// where the edges are straitgh lines connecting the given ordered set
// of points. The variable points is a list of point coordinates, each
// point coordinates is a tuple (x,y). If closed is true, then and edge
// is added to connect the last point to the first, and then create a
// closed polyline, i.e. a polygone
func (s *Sketcher) Polyline(points []struct{ X, Y float64 }, closed bool) {
	p := points[0]
	s.MoveTo(p.X, p.Y)
	for _, p = range points[1:] {
		s.LineTo(p.X, p.Y)
	}
	if closed {
		p = points[0]
		s.LineTo(p.X, p.Y)
	}
	s.x = p.X
	s.y = p.Y
}

// --------------------------------------------------------------------
// Write text functions

func (s *Sketcher) Text(x, y float64, text string) {
	px, py := s.canvasCoordinates(x, y)
	style := s.Pencil.TextStyle()
	s.body += fmt.Sprintf(textPattern, px, py, style, text) + "\n"
}

func (s *Sketcher) PointWithLabel(x, y float64, label string) {
	xoffset := 0.01 // suppose an x range = [0,1]
	yoffset := 0.01 // suppose an y range = [0,1]
	s.Point(x, y)
	s.Text(x+xoffset, y+yoffset, label)
}
