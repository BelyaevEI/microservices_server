package main

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/BelyaevEI/microservices_chat/internal/app"
	desc "github.com/BelyaevEI/microservices_chat/pkg/chat_v1"
	sq "github.com/Masterminds/squirrel"
)

var configPath string

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func init() {
	configPath = os.Getenv("CONFIG_PATH")
}

// Create chat
func (s *server) CreateChat(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	if req.GetChatname() == "" {
		return nil, status.Error(codes.InvalidArgument, "chat name is empty")
	}

	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "user_ids").
		Values(req.GetChatname(), req.GetId()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("create chat build is failed: %e", err)
		return nil, status.Error(codes.Internal, "create chat build is failed")
	}

	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Printf("err: %e", err)
		return nil, status.Error(codes.Internal, "failed to insert table chats")
	}

	return &desc.CreateResponse{
		Id: chatID,
	}, nil
}

// Delete chat
func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	builderDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("delete chat build is failed: %v", err)
		return nil, status.Error(codes.Internal, "failed to build query")
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete chat: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete chat")
	}

	log.Printf("deleted %d rows", res.RowsAffected())
	return nil, nil
}

// Send message to server
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {

	if req.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "message is empty")
	}

	builderInsert := sq.Insert("message").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text").
		Values(req.GetToChatId(), req.GetFromUserId(), req.GetText()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("err: %e", err)
		return nil, status.Error(codes.Internal, "send message failed to build query")
	}
	var messageID uuid.UUID
	err = s.pool.QueryRow(ctx, query, args...).Scan(&messageID)
	if err != nil {
		log.Printf("err: %e", err)
		return nil, status.Error(codes.Internal, "send message failed to insert message")
	}

	log.Printf("inserted chat with id: %d", messageID)

	return &desc.SendMessageResponse{
		Id:     messageID.String(),
		ChatId: req.GetToChatId(),
	}, nil
}

func main() {

	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

}
