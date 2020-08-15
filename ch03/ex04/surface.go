package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            //キャンパスの大きさ
	cells         = 100                 //升目の数
	xyrange       = 30.0                //軸の範囲
	xyscale       = width / 2 / xyrange //x単位、y単位当たりの画素数
	zscale        = height * 0.4        //z単位当たりの画素数
	angle         = math.Pi / 6         //x,y軸の角度

	maxHeight = 1.0
	minHeight = -0.22

	red  = 0x00ff0000
	blue = 0x000000ff
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30°),cos(30°)

func main() {
	http.HandleFunc("/", draw)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func draw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ah, ok := corner(i+1, j)
			bx, by, bh, ok := corner(i, j)
			cx, cy, ch, ok := corner(i, j+1)
			dx, dy, dh, ok := corner(i+1, j+1)

			if !ok {
				println("return f() is Isf or NaN!!")
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, color(ah, bh, ch, dh))
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ok := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if !ok {
		return sx, sy, z, false
	}
	return sx, sy, z, true

}

func f(x, y float64) (value float64, ok bool) {

	r := math.Hypot(x, y)
	value = math.Sin(r) / r

	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, false
	}
	return value, true
}

func color(ah, bh, ch, dh float64) string {
	height := (ah + bh + ch + dh) / 4

	if height > maxHeight || minHeight > height {
		panic(fmt.Sprintf("illegal height : %g", height))
	}

	delta := uint32((maxHeight - height) / (maxHeight - minHeight) * 255)

	color := prependZeros(fmt.Sprintf("#%X", (red-delta<<16)+delta))

	return color
}

func prependZeros(hex string) string {
	result := hex
	for i := len(hex); i < 6; i++ {
		result = "0" + result
	}
	return result
}
