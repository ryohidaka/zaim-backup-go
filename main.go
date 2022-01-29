package main

import "fmt"

func main() {

	// クライアント設定
	c := GetClient()

	// データ取得
	d := GetZaimData(c)
	fmt.Println(d)

}
