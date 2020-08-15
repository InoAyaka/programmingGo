package main

import (
	"flag"
	"fmt"
	"os"
)

var unit = flag.String("u", "", "unit")
var value = flag.Int("v", 0, "value")

func main() {
	flag.Parse()

	switch *unit {
	case "length":
		valueCheck()
		lengthConv()
	case "wight":
		valueCheck()
		wightConv()
	default:
		fmt.Fprintln(os.Stderr, "変換する単位を選択してください。（length or wight）")
	}
}

func valueCheck() {
	if *value == 0 {
		fmt.Println("変換する値を入力してください")
		_, err := fmt.Scan(value)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	fmt.Println(*value, "を変換します")
}

//長さ
type feet float64
type meters float64

func lengthConv() {
	fmt.Printf("%v メートル　＝　%g フィート\n", *value, mToF())
	fmt.Printf("%v フィート　＝　%g メートル\n", *value, fToM())
}

func mToF() feet {
	//1メートル ＝ 3.28084フィート
	return (feet)((float64)(*value) * 3.28084)
}

func fToM() meters {
	return (meters)((float64)(*value) / 3.28084)
}

//重さ
type lb float64
type kg float64

func wightConv() {
	fmt.Printf("%v ポンド　＝　%g キログラム\n", *value, lbToKg())
	fmt.Printf("%v キログラム　＝　%g ポンド\n", *value, kgToLb())
}

func lbToKg() kg {
	//1lb = 0.45392kg
	return (kg)((float64)(*value) * 0.45392)
}

func kgToLb() lb {
	return (lb)((float64)(*value) / 0.45392)
}
