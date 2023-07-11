package request

type NotificationRequest struct {
	Flame       bool    `json:"flame"`
	Temperature float64 `json:"temperature"`
	Gas         int     `json:"gas"`
	Message     string  `json:"message"`
	State       int     `json:"state"`
}

type SendMessageTelegramRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}
