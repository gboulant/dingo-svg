package svg

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
