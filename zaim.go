package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/gocarina/gocsv"
	gozaim "github.com/s-sasaki-0529/go-zaim"
)

type ZaimData struct {
	money      []gozaim.Money
	categories []gozaim.Category
	genres     []gozaim.Genre
	accounts   []gozaim.Account
}

type MoneyJP struct {
	Date         string `csv:"日付"`
	Mode         string `csv:"方法"`
	Category     string `csv:"カテゴリ"`
	Genre        string `csv:"カテゴリの内訳"`
	From         string `csv:"支払元"`
	To           string `csv:"入金先"`
	Name         string `csv:"品目"`
	Comment      string `csv:"メモ"`
	Place        string `csv:"お店"`
	CurrencyCode string `csv:"通貨"`
	Income       int    `csv:"収入"`
	Payment      int    `csv:"支出"`
	Transfer     int    `csv:"振替"`
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

// 種別IDをもとに、種別名を付与する
func ConvertData(datas ZaimData) []MoneyJP {
	var money []MoneyJP

	for _, v := range datas.money {
		p, i, t := GetAmount(v.Mode, v.Amount)
		cm := MoneyJP{
			v.Date,
			v.Mode,
			GetCategoryName(v.CategoryID, datas.categories),
			GetGenreName(v.GenreID, datas.genres),
			GetAccountName(v.FromAccountID, datas.accounts),
			GetAccountName(v.ToAccountID, datas.accounts),
			v.Name,
			v.Comment,
			v.Place,
			v.CurrencyCode,
			p,
			i,
			t,
		}

		money = append(money, cm)

	}
	return money
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

// IDに紐づくジャンル名を返却
func GetGenreName(id int, genres []gozaim.Genre) string {
	for _, v := range genres {
		if v.ID == id {
			return v.Name
		}
	}
	return ""
}

// IDに紐づく口座名を返却
func GetAccountName(id int, accounts []gozaim.Account) string {
	for _, v := range accounts {
		if v.ID == id {
			return v.Name
		}
	}
	return "-"
}

// 方法に応じた額を返却
func GetAmount(mode string, amount int) (int, int, int) {
	var p, i, t int

	switch mode {
	case "payment":
		p = amount
	case "income":
		i = amount
	case "transfer":
		t = amount
	default:
	}

	return p, i, t

}

// JSON出力
func OutputJSON(money gozaim.MoneySlice) {
	file, err := json.MarshalIndent(money, "", " ")
	if err != nil {
		fmt.Println("Failed to convert json", err)
	}

	err = ioutil.WriteFile("zaim-backup.json", file, 0644)
	if err != nil {
		fmt.Println("Failed to write json", err)
	}

}

// CSV出力
func OutputCSV(money []MoneyJP) {
	file, err := os.OpenFile("zaim-backup.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to write csv", err)
	}
	defer file.Close()

	// csvファイルを書き出し
	gocsv.MarshalFile(&money, file)

}
