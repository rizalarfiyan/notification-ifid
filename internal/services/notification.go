package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rizalarfiyan/notification-ifid/config"
	"github.com/rizalarfiyan/notification-ifid/constants"
	"github.com/rizalarfiyan/notification-ifid/internal/request"
	"github.com/rizalarfiyan/notification-ifid/utils"
)

type NotificationService interface {
	SendNotification(params request.NotificationRequest) error
}

type notificationService struct{}

func NewNotificationService() NotificationService {
	return &notificationService{}
}

func (h *notificationService) SendNotification(params request.NotificationRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conf := config.Get()
	bodyParams := request.SendMessageTelegramRequest{
		ChatID:    conf.TelegramChatId,
		Text:      h.parseToMarkdownText(params),
		ParseMode: "Markdown",
	}

	byteData, err := json.Marshal(bodyParams)
	if err != nil {
		return err
	}

	byteReader := bytes.NewReader(byteData)
	url := fmt.Sprint(constants.TELEGRAM_API, conf.TelegramToken, "/sendMessage")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, byteReader)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (h *notificationService) parseToMarkdownText(params request.NotificationRequest) string {
	format := `Informasi Terbaru! - *%s*
*Message*  : *%s*
*Temp*        : %.2f °C
*Gas*           : %d PPM
*Flame*       : %t`
	return fmt.Sprintf(format, utils.GetStateWording(params.State), params.Message, params.Temperature, params.Gas, params.Flame)
}
