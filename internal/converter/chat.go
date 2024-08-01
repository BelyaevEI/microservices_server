package converter

import (
	"github.com/BelyaevEI/microservices_chat/internal/model"
	desc "github.com/BelyaevEI/microservices_chat/pkg/chat_v1"
)

// ToChatCreateFromDesc converts desc.ChatCreate to model.ChatCreate
func ToChatCreateFromDesc(chatCreate *desc.CreateRequest) *model.ChatCreate {
	return &model.ChatCreate{
		Name:   chatCreate.Chatname,
		UserID: chatCreate.Id,
	}
}

// ToMessageCreateFromDesc converts desc.MessageCreate to model.MessageCreate
func ToMessageCreateFromDesc(messageCreate *desc.SendMessageRequest) *model.MessageCreate {
	return &model.MessageCreate{
		Info: model.MessageInfo{
			ChatID: messageCreate.ToChatId,
			UserID: messageCreate.FromUserId,
			Text:   messageCreate.Text,
		},
	}
}
