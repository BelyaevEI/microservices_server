package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/model"
)

func (s *serv) CreateChat(ctx context.Context, createChat *model.ChatCreate) (int64, error) {
	var id int64

	id, err := s.chatRepository.CreateChat(ctx, createChat)
	if err != nil {
		return 0, err
	}

	return id, nil
}
