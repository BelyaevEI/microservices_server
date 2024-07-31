package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/model"
)

func (s *serv) CreateChat(ctx context.Context, createChat *model.ChatCreate) (int64, error) {
	return 0, nil
}
