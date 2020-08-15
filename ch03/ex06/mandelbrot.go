package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, xmax = -2, +2
	ymin, ymax = -2, +2
	width      = 1024
	height     = 1024
)

func main() {

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

	superSample(img)

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

func superSample(img *image.RGBA) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r1, g1, b1, a1 := img.At(x, y).RGBA()
			r2, g2, b2, a2 := img.At(x+1, y).RGBA()
			r3, g3, b3, a3 := img.At(x, y+1).RGBA()
			r4, g4, b4, a4 := img.At(x+1, y+1).RGBA()

			r, g, b, a := (r1+r2+r3+r4)/4, (g1+g2+g3+g4)/4, (b1+b2+b3+b4)/4, (a1+a2+a3+a4)/4

			img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})

		}
	}
}
