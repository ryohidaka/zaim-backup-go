package main

import (
	"fmt"
	"os"

	gozaim "github.com/s-sasaki-0529/go-zaim"

	"github.com/joho/godotenv"
)

func GetClient() *gozaim.Client {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	c := gozaim.NewClient(
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"),
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_SECRET"),
	)

	return c
}
