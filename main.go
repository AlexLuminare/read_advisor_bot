package main

import (
	"fmt"
	"github.com/AlexLuminare/read_advisor_bot/clients/telegram"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	//полачем данные из .env
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	token := os.Getenv("TG_TOKEN")
	host := os.Getenv("HOST")
	fmt.Printf("token: %s", token)

	//
	tgClient := telegram.New(host, token)

	//fetcher =fetcher.New()

	//processor =processor.New()

	//consumer.Start(fetcher, processor)

}
