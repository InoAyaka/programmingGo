package main

func main() {
	//検証：nilマップへの代入で静的解析に怒られるか
	var m map[int]string

	m[1] = "aaa"

}
