package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            //キャンパスの大きさ
	cells         = 100                 //升目の数
	xyrange       = 30.0                //軸の範囲
	xyscale       = width / 2 / xyrange //x単位、y単位当たりの画素数
	zscale        = height * 0.4        //z単位当たりの画素数
	angle         = math.Pi / 6         //x,y軸の角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30°),cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			bx, by, ok := corner(i, j)
			cx, cy, ok := corner(i, j+1)
			dx, dy, ok := corner(i+1, j+1)

			if !ok {
				println("return f() is Isf or NaN!!")
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ok := f(x, y)

	if !ok {
		return 0.0, 0.0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true

}

func f(x, y float64) (value float64, ok bool) {

	r := math.Hypot(x, y)
	value = math.Sin(r) / r

	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, false
	}
	return value, true
}
