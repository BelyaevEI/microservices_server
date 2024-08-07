package chat

import (
	"github.com/BelyaevEI/microservices_chat/internal/service"
	desc "github.com/BelyaevEI/microservices_chat/pkg/chat_v1"
)

// Implementation represents a chat API implementation.
type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewImplementation creates a new chat API implementation.
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
