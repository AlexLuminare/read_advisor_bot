package telegram

import "github.com/AlexLuminare/read_advisor_bot/clients/telegram"

// Реализует интерфейсы Fetcher и Processor
type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
}

//func NewProcessor(tg *telegram.Client, storage)
