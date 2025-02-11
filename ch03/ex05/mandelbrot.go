package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, xmax = -2, +2
		ymin, ymax = -2, +2
		width      = 1024
		height     = 1024
	)

	//指定された境界（今回はRectangle:長方形）を持つRGBAを返す
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			//z = x + yi
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)

}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	blue := color.RGBA{0, 0, 255, 255}

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 0, 255 - contrast*n, 255}
		}
	}
	return blue
}
