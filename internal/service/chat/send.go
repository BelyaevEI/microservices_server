package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error) {
	return "", nil
}
