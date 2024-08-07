package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/client/postgres"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) DeleteChat(ctx context.Context, id int64) error {

	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := postgres.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}
