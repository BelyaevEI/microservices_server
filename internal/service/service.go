package service

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/model"
)

type ChatService interface {
	CreateChat(ctx context.Context, createChat *model.ChatCreate) (int64, error)
	SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error)
	DeleteChat(ctx context.Context, id int64) error
}
