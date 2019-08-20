package main

import (
	"image"
	"image/color"
	"math"
)

const (
	// W width
	W int = 512
	// H height
	H          int = 512
	whiteIndex     = 0 // first color in palette
	blackIndex     = 1 // next color in palette
)

var palette = []color.Color{color.White, color.Black}

func setPixel(x, y int, img image.Image) {
}

func ternary(expr bool, tval, fval int) int {
	if expr {
		return tval
	}
	return fval
}

func bresenham(x0, y0, x1, y1 int, setPixel func(int, int)) {
	var sx, sy, e1 int
	dx := int(math.Abs(float64(x1 - x0)))
	dy := int(math.Abs(float64(y1 - y0)))

	sx = ternary(x0 < x1, 1, -1)
	sy = ternary(y0 < y1, 1, -1)
	e1 = ternary(dx > dy, dx/2, -dy/2)

	for ; x0 != x1 || y0 != y1; setPixel(x0, y0) {
		e2 := e1
		if e2 > -dx {
			e1 -= dy
			x0 += sx
		}
		if e2 < dy {
			e1 += dx
			y0 += sy
		}
	}
}

func main() {

}
