package chat

import (
	"github.com/BelyaevEI/microservices_chat/internal/client/postgres"
	"github.com/BelyaevEI/microservices_chat/internal/repository"
)

const (
	tableName  = "chat"
	idColumn   = "id"
	nameColumn = "name"
	user_ids   = "user_ids"

	tableNameMessage = "message"
	chatIDColumn     = "chat_id"
	userIDColumn     = "user_id"
	textColumn       = "text"
)

type repo struct {
	db postgres.Client
}

// NewRepository creates a new user repository.
func NewRepository(db postgres.Client) repository.ChatRepository {
	return &repo{db: db}
}
