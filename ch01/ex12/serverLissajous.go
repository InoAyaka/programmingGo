// 検証　p.25 GIFアニメーションをHTTPクライアントへ書き込み
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//パラメータの取得と変換
	cycles := getCycle(r.URL.String())
	lissajous(w, cycles)
}

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

func lissajous(out io.Writer, cycles int) {
	const (
		//cycle   = 5     //発振器xが完了する周回の回数
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
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

//URLのパラメータから、サイクルを取得、intに変換
func getCycle(uS string) int {
	//URLを解析
	u, err := url.Parse(uS)

	if err != nil {
		log.Fatal(err)
	}

	//パラメータを取得
	m := u.Query()

	//デフォルトのサイクルを5に設定
	cycles := 5

	//cycles に設定された値を取得し（複数の場合には最初のもののみ）、intに変換
	for key, vals := range m {
		if key == "cycles" {
			for _, val := range vals {
				if i, err := strconv.Atoi(val); err != nil {
					log.Fatal(err)
				} else {
					cycles = i
				}
			}
		}
	}

	return cycles

}
