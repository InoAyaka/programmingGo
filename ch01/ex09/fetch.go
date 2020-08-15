package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		//URLのチェック
		url = checkURL(url)

		//HTTPリクエストの送信とレスポンスの受信
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//レスポンスのBobyを読み込み
		//b, err := ioutil.ReadAll(resp.Body)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//リークを防ぐためにクローズ
		resp.Body.Close()

		fmt.Println("HTTP status : ", resp.Status)

	}
}

//引数のチェック（頭に"http://"があるか）
func checkURL(url string) string {

	if !strings.HasPrefix(url, "http://") {
		return "http://" + url
	}
	return url
}
