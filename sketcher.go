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
// Coordinates System
// ===========================================================================

type CoordinateSystem struct {
	cnvxsize int
	cnvysize int
}

func NewCoordinateSystem() *CoordinateSystem {
	return &CoordinateSystem{600, 600}
}

func (c CoordinateSystem) canvasCoordinates(x, y float64) (px, py float64) {
	// WRN: x and y are supposed to be values between 0 and 1. They are mapped to pixels
	px = x * float64(c.cnvxsize)
	py = float64(c.cnvysize) * (1 - y)
	return px, py
}

func (c CoordinateSystem) canvasScaling(size float64) (psize float64) {
	psize = float64(c.cnvxsize) * size
	return psize
}

// ===========================================================================
// Pencil and style management
// ===========================================================================
const (
	drawStylePattern = "stroke: %s; stroke-width: %d; fill: %s"
	textStylePattern = "font-family:%s; font-size:%d; font-weight:%s; fill: %s"
)

const (
	// Constant for the drawing
	defaultLineColor = "black"
	defaultLineWidth = 2
	defaultFillColor = defaultLineColor
	// Constant for the text
	defaultFontFamily = "Arial"
	defaultFontWeight = "normal"
	defaultFontSize   = 20
	defaultFontColor  = defaultLineColor
)

type Pencil struct {
	// Parameters for the drawing
	LineColor string
	LineWidth int
	FillColor string
	FillMode  bool // if true, fill any closed shape with the FillColor color

	// Parameters for the text
	FontFamily string
	FontWeight string
	FontSize   int
	FontColor  string
}

func NewPencil(linecolor string, linewidth int) *Pencil {
	return &Pencil{
		LineColor:  linecolor,
		LineWidth:  linewidth,
		FillColor:  linecolor,
		FillMode:   true,
		FontFamily: defaultFontFamily,
		FontWeight: defaultFontWeight,
		FontSize:   defaultFontSize,
		FontColor:  linecolor,
	}
}

func (p Pencil) DrawStyleWithFillMode(fill bool) string {
	fillcolor := p.FillColor
	if !fill {
		fillcolor = "none"
	}
	return fmt.Sprintf(drawStylePattern, p.LineColor, p.LineWidth, fillcolor)
}

func (p Pencil) DrawStyle() string {
	return p.DrawStyleWithFillMode(p.FillMode)
}

func (p Pencil) TextStyle() string {
	return fmt.Sprintf(textStylePattern, p.FontFamily, p.FontSize, p.FontWeight, p.FontColor)
}

var defaultPencil *Pencil = NewPencil(defaultLineColor, defaultLineWidth)

// ===========================================================================
// Sketch builder
// ===========================================================================

type Sketcher struct {
	x, y   float64
	body   string
	cs     *CoordinateSystem
	Pencil *Pencil
}

func NewSketcher() *Sketcher {
	cs := NewCoordinateSystem()
	return &Sketcher{x: 0., y: 0, body: "", cs: cs, Pencil: defaultPencil}
}

// --------------------------------------------------------------------
// Sketch management functions

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
	file, err := os.OpenFile(svgpath, os.O_CREATE|os.O_WRONLY, 0600)
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
// Drawing functions

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

func (s *Sketcher) Circle(cx, cy, r float64, fill bool) {
	pcx, pcy := s.canvasCoordinates(cx, cy)
	pr := s.canvasScaling(r)
	style := s.Pencil.DrawStyleWithFillMode(fill)
	s.body += fmt.Sprintf(circPattern, pcx, pcy, pr, style) + "\n"
	s.x = cx
	s.y = cy
}

func (s *Sketcher) Point(x, y float64) {
	r := s.pointSize()
	s.Circle(x, y, r, true)
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

func (s *Sketcher) Edge(x1, y1, x2, y2 float64) {
	s.MoveTo(x1, y1)
	s.LineTo(x2, y2)
}

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
