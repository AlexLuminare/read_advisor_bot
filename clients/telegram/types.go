package telegram

type UpdatesResponce struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID     int64  `json:"update_id"`
	Mesage string `json:"message"`
}
