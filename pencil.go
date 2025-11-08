package svg

import "fmt"

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
	defaultFillColor = "black"
	defaultFillMode  = true
	// Constant for the text
	defaultFontFamily = "Arial"
	defaultFontWeight = "normal"
	defaultFontSize   = 20
	defaultFontColor  = "black"
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
		FillColor:  defaultFillColor,
		FillMode:   defaultFillMode,
		FontFamily: defaultFontFamily,
		FontWeight: defaultFontWeight,
		FontSize:   defaultFontSize,
		FontColor:  defaultFontColor,
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
