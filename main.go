package main

import "fmt"

func main() {

	// クライアント設定
	c := GetClient()

	// データ取得
	d := GetZaimData(c)

	// 種別名を付与する
	cd := ConvertData(d)
	fmt.Println(cd)

}
