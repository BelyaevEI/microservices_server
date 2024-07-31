package chat

import (
	"github.com/BelyaevEI/microservices_chat/internal/client/postgres"
	"github.com/BelyaevEI/microservices_chat/internal/repository"
	"github.com/BelyaevEI/microservices_chat/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      postgres.TxManager
}

func NewService(chatRepository repository.ChatRepository, txManager postgres.TxManager) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
