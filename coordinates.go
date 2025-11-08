package svg

import (
	"fmt"
	"math"
)

// ===========================================================================
// Coordinates System
// ===========================================================================

const (
	DefaultCanvasWidth  = 600
	DefaultCanvasHeight = 600
)

type CoordinateSystem struct {
	cnvxsize   int     // canvas width in pixels
	cnvysize   int     // canvas height in pixels
	xorigin    float64 // position of the origin on the X axis in pixels
	yorigin    float64 // position of the origin on the Y axis in pixels
	xsign      float64 // orientation of the X axis: -1 means from right to left
	ysign      float64 // orientation of the Y axis: -1 means from bottom to top
	unit2pixel float64 // number of pixels in a user unit (could be a float)
}

// NewCoordinateSystem returns the default coordinates system. It is a
// coordinates system 1/ based on a canvas Width x Height, 2/ with origin at the
// bottom left corner, 3/ with user y axis oriented bottom up (inverse of the
// canvas native axis), and 4/ with an xrange equal to 1, i.e. x=1 corresponds
// to a x coordinates at the right boundary of the canvas (x=1 => xpixels =
// cnvwidth)
func NewCoordinateSystem() *CoordinateSystem {
	cnvxsize := DefaultCanvasWidth
	cnvysize := DefaultCanvasHeight
	xrange := 1. // coordinates are supposed to be in [0..1]
	return NewCoordSysBottomLeft(cnvxsize, cnvysize, xrange)
}

// canvasCoordinates returns the position of the point in the canvas native
// coordinates system, i.e. number of pixels from top left corner along the
// horizontal and vertical axis (oriented to the bottom) respectivelly
func (c CoordinateSystem) canvasCoordinates(x, y float64) (px, py float64) {
	px = c.xorigin + c.xsign*c.unit2pixel*x
	py = c.yorigin + c.ysign*c.unit2pixel*y
	return px, py
}

func (c CoordinateSystem) canvasScaling(size float64) (psize float64) {
	psize = c.unit2pixel * size
	return psize
}

// o2s return the sign to consider for the user axis orientation depending if
// the user axis is oriented as the canvas native axis (inverse = false) or
// inverse of the canvas native axis (inverse = true). If inverse is true, then
// it returns -1, else it return +1
func o2s(inverse bool) (sign float64) {
	if inverse {
		sign = -1
	} else {
		sign = +1
	}
	return sign
}

func newCoordSystemAtOrigin(cnvXorigin, cnvYorigin float64, cnvwidth, cnvheight int, xrange float64) *CoordinateSystem {
	var xinverse bool = false // oriented as the canvas native orientation
	var yinverse bool = true  // inverse of the canvas native orientation
	xsign := o2s(xinverse)
	ysign := o2s(yinverse)

	// we consider a user x axis form 0 to xrange, then the unit has a size in
	// pixels equal to the canvas width divided by the xrange
	unit2pixel := float64(cnvwidth) / xrange
	// As a consequence, x and y are supposed to be values between 0 and xrange,
	// otherwize, they will not be viewed inside the canvas boundaries

	return &CoordinateSystem{
		cnvxsize:   cnvwidth,
		cnvysize:   cnvheight,
		xorigin:    cnvXorigin,
		yorigin:    cnvYorigin,
		xsign:      xsign,
		ysign:      ysign,
		unit2pixel: unit2pixel,
	}
}

// NewCoordSysBottomLeft creates a Coordinate System whose origin is at the
// bottom left corner of the canvas, with y coordinates axis oriented bottom up
// (inverse of the canvas native Y axis). The xrange is the range of x values
// (xmax - xmin) from the left boundary of the canvas to the right boundary.
func NewCoordSysBottomLeft(cnvwidth, cnvheight int, xrange float64) *CoordinateSystem {
	cnvXorigin := 0.
	cnvYorigin := float64(cnvheight)
	return newCoordSystemAtOrigin(cnvXorigin, cnvYorigin, cnvwidth, cnvheight, xrange)
}

// NewCoordSysCentered creates a Coordinate System whose origin is at the center
// point of the canvas, with y coordinates axis oriented bottom up (inverse of
// the canvas native Y axis). The xrange is the range of x values (xmax - xmin)
// from the left boundary of the canvas to the right boundary.
func NewCoordSysCentered(cnvwidth, cnvheight int, xrange float64) *CoordinateSystem {
	cnvXorigin := float64(cnvwidth) * 0.5
	cnvYorigin := float64(cnvheight) * 0.5
	return newCoordSystemAtOrigin(cnvXorigin, cnvYorigin, cnvwidth, cnvheight, xrange)
}

func NewCoordSysBoundedBy(points []struct{ X, Y float64 }, cnvwidth, cnvheight int) *CoordinateSystem {
	xmin, ymin, xmax, ymax := boundingBox(points)
	fmt.Println(xmin, ymin, xmax, ymax)
	return nil
}

func boundingBox(points []struct{ X, Y float64 }) (xmin, ymin, xmax, ymax float64) {
	xmin = math.Inf(+1)
	xmax = math.Inf(-1)
	ymin = math.Inf(+1)
	ymax = math.Inf(-1)
	for _, p := range points {
		if p.X < xmin {
			xmin = p.X
		}
		if p.X > xmax {
			xmax = p.X
		}
		if p.Y < ymin {
			ymin = p.Y
		}
		if p.Y > ymax {
			ymax = p.Y
		}
	}
	return xmin, ymin, xmax, ymax
}
