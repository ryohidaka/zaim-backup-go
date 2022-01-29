package main

import (
	"fmt"
	"net/url"

	gozaim "github.com/s-sasaki-0529/go-zaim"
)

type ZaimData struct {
	money []gozaim.Money
}

// Zaimのデータを取得する
func GetZaimData(c *gozaim.Client) ZaimData {

	// データ一覧の取得
	m, err := c.FetchMoney(url.Values{})
	if err != nil {
		fmt.Println("Failed to get money", err)
	}

	msg := fmt.Sprintf("%d 件のデータを取得しました。\n", len(m))
	fmt.Println(msg)

	result := ZaimData{
		m,
	}

	return result
}
