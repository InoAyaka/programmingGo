package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

//色の指定
var palette = []color.Color{
	color.White,
	color.RGBA{0x00, 0x80, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0x80, 0xff},
}

const (
	whilteIndex = 0
	secondIndex = 1
	therdIndex  = 2
)

func main() {
	//乱数の作成のために、元になる種をセット（このコードでは現時刻を使用）
	rand.Seed(time.Now().UTC().UnixNano())
	//リサジュー図形の作成
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycle   = 5     //発振器xが完了する周回の回数
		res     = 0.001 //回転の分解能
		size    = 100   //画像キャンパス
		nframes = 64    //アニメーションフレーム数
		delay   = 8     //10ms単位のフレーム間の遅延
	)

	freq := rand.Float64() * 3.0 //発振器yの相対周波数

	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //位相差

	//アニメーションフレームの数分ループ
	for i := 0; i < nframes; i++ {
		//作成する画像のサイズ、色の定義
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		var index uint8

		if i > 32 {
			index = secondIndex
		} else {
			index = therdIndex
		}

		//5回円を描くまで続ける
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	//GIFアニメを生成
	gif.EncodeAll(out, &anim)

}
