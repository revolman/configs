package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
)

var palette color.Color

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное число z.
			img.Set(px, py, mandelbrot(z, r, g, b))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128, r int, g int, b int) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			palette = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(contrast * n)}
			return palette
		}
	}
	//	palette = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	return color.Black
}
