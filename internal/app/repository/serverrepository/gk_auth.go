package serverrepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (r *KeeperRepository) AddUser(userID, login, password string) error {
	insertTag, err := r.db.Exec(context.Background(), "INSERT INTO users (user_id, login, password) values($1, $2, $3)", userID, login, password)
	if err != nil {
		return fmt.Errorf("insert in users table failed with error: %w", err)
	}

	if insertTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected: %w", pgx.ErrNoRows)
	}

	return nil
}

func (r *KeeperRepository) CheckUser(login, password string) (string, error) {
	var userID string
	err := r.db.QueryRow(context.Background(), "SELECT user_id FROM users WHERE login=$1 AND password=$2", login, password).Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("select user from users failed with error: %w", ErrUserNotFound)
		}
		return "", fmt.Errorf("select user from users failed with error: %w", err)
	}

	return userID, nil
}
