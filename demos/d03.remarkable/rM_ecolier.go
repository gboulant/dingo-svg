package main

import svg "github.com/gboulant/dingo-svg"

type COLORS_MAP map[string]string

var COLORS_CATALOG map[string]COLORS_MAP = map[string]COLORS_MAP{
	"original": {
		"blue": "blue",
		"red":  "red",
	},
	// See conversion at https://www.orbworks.com/cedocs/color.htm
	"grayscale": {
		"blue": "Silver",
		"red":  "Gray",
	},
}

func rM_ecolier(svgpath string) error {
	s := NewRemarkableSketcher()
	xmin, xmax, ymin, ymax := s.CoordinatesSystem().UserCoordinatesBoundaries()

	colors := COLORS_CATALOG["original"]
	p_redbold := svg.NewPencil(colors["red"], 3)
	p_bluethin := svg.NewPencil(colors["blue"], 1)
	p_bluebold := svg.NewPencil(colors["blue"], 3)

	wmargin := 34.
	hheader := 20. // mm
	hfooter := 15.

	nbticsByCell := 4
	cellsize := 8. // mm.
	ticsize := float64(cellsize) / float64(nbticsByCell)

	// The page height is composed of:
	//
	// - a header of size hheader
	// - nbcells cells of size cellsize (height) where nbcells is an integer
	// - a footer of size hfooter
	//
	// And the total height of the page H should be the sum of these quantities,
	// which means that the number of cells nbcells should be:
	//
	//   nbcells = (H - hheader - hfooter) / cellsize
	//
	// But, if we choose hheader and hfooter arbitrarily (H and cellsize are
	// fixed), then in most of cases nbcells would be a real, not and integer.
	// So, we consider nbcells as the integer part of the result, and consider the
	// decimal part as a fraction of cellsize to add the either the footer or the
	// header.
	//
	// Moreover, the cells lines does not start at the footer or the header,
	// there is a number of tics nbTicsHeader above the first line (between the
	// header and the first line) and a number of tics nbTicsFooter after the
	// last line (before the footer).
	nbTicsHeader := 3 // nb of tics (interlines) above the first line
	nbTicsFooter := 2 // nb of tics (interlines) after the last line
	//
	// Then the equation to solve is in fact:
	//
	//   H = nbcells*cellsize + (nbTicsHeader+nbTicsFooter)*ticsize + hheader + hfooter
	//
	// with ticsize = cellsize / nbticsByCell, then:
	//
	//   H = nbcells*cellsize + (nbTicsHeader+nbTicsFooter)*cellsize/nbticsByCell + hheader + hfooter
	//
	// Then:
	//
	//   nbcells = (H - hheader - hfooter)/cellsize - (nbTicsHeader+nbTicsFooter)/nbticsByCell

	H := ymax - ymin
	nbcellsfloat := float64(H-hheader-hfooter)/cellsize - float64(nbTicsHeader+nbTicsFooter)/float64(nbticsByCell)
	nbcells := int(nbcellsfloat)
	decimal := nbcellsfloat - float64(nbcells)

	// We add the remaining quantity (decimal part) to the footer
	hfooter += decimal * cellsize

	// --- red left margin ---------------------------------------
	s.Pencil = p_redbold
	x := xmin + wmargin
	s.MoveTo(x, 0)
	s.LineTo(x, ymax)

	// --- thin blue horizontal lines (interlines) ---------------
	s.Pencil = p_bluethin
	y := ymax - hheader // start at the bottom
	for y >= hfooter {
		s.MoveTo(xmin, y)
		s.LineTo(xmax, y)
		y -= ticsize
	}

	// --- bold blue horizontal lines (lines) ---------------------
	s.Pencil = p_bluebold
	y = ymax - (hheader + float64(nbTicsHeader)*ticsize)
	for y >= hfooter+float64(nbTicsFooter)*ticsize {
		s.MoveTo(xmin, y)
		s.LineTo(xmax, y)
		y -= cellsize
	}

	// --- bold blue vertical lines -------------------------------
	x = xmin + wmargin + cellsize
	for x <= xmax {
		s.MoveTo(x, ymin)
		s.LineTo(x, ymax)
		x += cellsize
	}

	return s.Save(svgpath)
}
