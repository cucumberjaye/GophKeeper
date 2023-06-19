package serverrepository

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrUserNotFound = errors.New("user not found")                         // Ошибка отсутсвия пользователя в базе.
	ErrDataNotFound = errors.New("data not found")                         // Ошибка остутсвия данных в базе.
	ErrUpdateLate   = errors.New("data updated from another client early") // Ошибка, если данные были изменены ранее с другого клиента.
)

// KeeperRepository - структура содержит подключение к базе данных.
type KeeperRepository struct {
	db *pgxpool.Pool
}

// New - инициализирует структуру KeeperRepository.
func New(pool *pgxpool.Pool) *KeeperRepository {
	return &KeeperRepository{
		db: pool,
	}
}
