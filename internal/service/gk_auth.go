package service

import (
	"github.com/cucumberjaye/GophKeeper/internal/pkg/tokens"
	"github.com/cucumberjaye/GophKeeper/pkg/hasher"
	"github.com/google/uuid"
)

func (s *KeeperService) AddUser(login, password string) error {
	hashPassword := hasher.HasherSha256(password)
	userID := uuid.New().String()
	return s.repository.AddUser(userID, login, string(hashPassword))
}

func (s *KeeperService) CreateToken(login, password string) (string, error) {
	hashPassword := hasher.HasherSha256(password)
	userID, err := s.repository.CheckUser(login, hashPassword)
	if err != nil {
		return "", err
	}

	return tokens.CreateToken(userID)
}
