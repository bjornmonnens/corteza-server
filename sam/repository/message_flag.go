package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/titpetric/factory"

	"github.com/crusttech/crust/sam/types"
)

type (
	MessageFlagRepository interface {
		With(ctx context.Context, db *factory.DB) MessageFlagRepository

		FindByID(ID uint64) (*types.MessageFlag, error)
		FindByMessageIDs(IDs ...uint64) ([]*types.MessageFlag, error)
		FindByFlag(messageID, userID uint64, flag string) (*types.MessageFlag, error)
		Create(mod *types.MessageFlag) (*types.MessageFlag, error)
		DeleteByID(ID uint64) error
	}

	messageFlag struct {
		*repository
	}
)

const (
	ErrMessageFlagNotFound = repositoryError("MessageFlagNotFound")
)

func MessageFlag(ctx context.Context, db *factory.DB) MessageFlagRepository {
	return (&messageFlag{}).With(ctx, db)
}

func (r *messageFlag) With(ctx context.Context, db *factory.DB) MessageFlagRepository {
	return &messageFlag{
		repository: r.repository.With(ctx, db),
	}
}

func (r *messageFlag) FindByID(ID uint64) (*types.MessageFlag, error) {
	sql := "SELECT * FROM message_flags WHERE id = ?"
	mod := &types.MessageFlag{}
	return mod, isFound(r.db().Get(mod, sql, ID), mod.ID > 0, ErrMessageFlagNotFound)
}

func (r *messageFlag) FindByFlag(messageID, userID uint64, flag string) (*types.MessageFlag, error) {
	args := []interface{}{messageID, flag}
	sql := "SELECT * FROM message_flags WHERE rel_message = ? AND flag = ? "

	if userID > 0 {
		sql += "AND rel_user = ? "
		args = append(args, userID)
	}

	mod := &types.MessageFlag{}
	return mod, isFound(r.db().Get(mod, sql, args...), mod.ID > 0, ErrMessageFlagNotFound)
}

// FindByMessageRange returns all flags by message id range
func (r *messageFlag) FindByMessageIDs(IDs ...uint64) ([]*types.MessageFlag, error) {
	rval := make([]*types.MessageFlag, 0)

	sql := `SELECT * FROM message_flags WHERE rel_message IN (?)`

	if sql, args, err := sqlx.In(sql, IDs); err != nil {
		return nil, err
	} else {
		return rval, r.db().Select(&rval, sql, args...)
	}
}

func (r *messageFlag) Create(mod *types.MessageFlag) (*types.MessageFlag, error) {
	mod.ID = factory.Sonyflake.NextID()
	mod.CreatedAt = time.Now()
	return mod, r.db().Insert("message_flags", mod)
}

func (r *messageFlag) DeleteByID(ID uint64) error {
	return exec(r.db().Exec("DELETE FROM message_flags WHERE id = ?", ID))
}