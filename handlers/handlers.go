package handlers

import (
	"awesome-project/database"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB  *gorm.DB
	Bot *tgbotapi.BotAPI
}

func (h *Handler) Pong(update tgbotapi.Update) error {
	text := "pong"

	var user database.User
	if err := h.DB.First(&user, update.Message.From.ID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			h.DB.Save(&database.User{
				UserID:    update.Message.From.ID,
				FirstName: update.Message.From.FirstName,
			})

			text += "!"
		} else {
			return err
		}
	} else {
		text += fmt.Sprintf(", mr. %s!", user.FirstName)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := h.Bot.Send(msg); err != nil {
		return err
	}

	return nil
}
