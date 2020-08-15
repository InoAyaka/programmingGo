package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

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

	}
}
