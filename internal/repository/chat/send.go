package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/model"
)

func (r *repo) SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error) {
	return "", nil
}
