package serverrepository

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDataNotFound = errors.New("data not found")
)

type KeeperRepository struct {
	db *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *KeeperRepository {
	return &KeeperRepository{
		db: pool,
	}
}
