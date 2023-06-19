package service

import (
	"github.com/cucumberjaye/GophKeeper/internal/pkg/tokens"
	"github.com/cucumberjaye/GophKeeper/pkg/hasher"
	"github.com/google/uuid"
)

// AddUser - функция для хеширования пароля и генерации id для пользователя и передачи их в слой репозитория.
func (s *KeeperService) AddUser(login, password string) error {
	hashPassword := hasher.HasherSha256(password)
	userID := uuid.New().String()
	return s.repository.AddUser(userID, login, string(hashPassword))
}

// CreateToken - создает jwt токен для пользоваеля.
func (s *KeeperService) CreateToken(login, password string) (string, error) {
	hashPassword := hasher.HasherSha256(password)
	userID, err := s.repository.CheckUser(login, hashPassword)
	if err != nil {
		return "", err
	}

	return tokens.CreateToken(userID)
}
