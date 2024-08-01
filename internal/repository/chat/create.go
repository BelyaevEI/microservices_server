package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/client/postgres"
	"github.com/BelyaevEI/microservices_chat/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) CreateChat(ctx context.Context, createChat *model.ChatCreate) (int64, error) {
	var id int64

	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn).
		Values(createChat.Name).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := postgres.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
