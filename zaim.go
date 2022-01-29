package main

import (
	"fmt"
	"net/url"

	gozaim "github.com/s-sasaki-0529/go-zaim"
)

type ZaimData struct {
	money      []gozaim.Money
	categories []gozaim.Category
	genres     []gozaim.Genre
	accounts   []gozaim.Account
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

	// カテゴリ一覧取得
	ca, err := c.FetchCategories()
	if err != nil {
		fmt.Println("Failed to get categories", err)
	}

	// ジャンル一覧取得
	g, err := c.FetchGenres()
	if err != nil {
		fmt.Println("Failed to get genres", err)
	}

	// 口座一覧取得
	a, err := c.FetchAccounts()
	if err != nil {
		fmt.Println("Failed to get accounts", err)
	}

	result := ZaimData{
		m,
		ca,
		g,
		a,
	}

	return result
}

// IDに紐づくカテゴリ名を返却
func GetCategoryName(id int, categories []gozaim.Category) string {
	for _, v := range categories {
		if v.ID == id {
			return v.Name
		}
	}
	return ""
}
