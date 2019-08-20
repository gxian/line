package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	// W width
	W int = 512
	// H height
	H int = 512
	// index
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

var palette = []color.Color{color.White, color.Black}

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
	rect := image.Rect(0, 0, W, H)
	plt := image.NewPaletted(rect, palette)
	var cx, cy float64
	cx = float64(W)*0.5 - 0.5
	cy = float64(H)*0.5 - 0.5
	for j := 0; j < 5; j++ {
		var r1, r2 float64
		m := math.Min(float64(W), float64(H))
		fj := float64(j)
		r1 = m * (fj + 0.5) * 0.085
		r2 = m * (fj + 1.5) * 0.085
		t := fj * math.Pi / 64.0
		for i := 1; i <= 64; i++ {
			ct := math.Cos(t)
			st := math.Sin(t)
			bresenham(
				int(cx+r1*ct),
				int(cy-r1*st),
				int(cx+r2*ct),
				int(cy-r2*st),
				func(x, y int) {
					plt.SetColorIndex(x, y, blackIndex)
				})
			t += 2.0 * math.Pi / 64.0
		}
	}

	f, err := os.Create("line_bresenham.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, plt); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
