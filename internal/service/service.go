package service

import "github.com/cucumberjaye/GophKeeper/internal/models"

type KeeperRepository interface {
	AuthRepository
	StoreRepository
}

type AuthRepository interface {
	AddUser(userID, login, password string) error
	CheckUser(login, password string) (string, error)
}

type StoreRepository interface {
	SetOrUpdateLoginPasswordData(userID string, data models.LoginPasswordData) error
	SetOrUpdateTextData(userID string, data models.TextData) error
	SetOrUpdateBinaryData(userID string, data models.BinaryData) error
	SetOrUpdateBankCardData(userID string, data models.BankCardData) error
	GetData(key, userID string) (any, error)
	GetDataArray(userID string) ([]string, error)
	DeleteData(key, userID string) error
}

type KeeperService struct {
	repository KeeperRepository
}

func New(repo KeeperRepository) *KeeperService {
	return &KeeperService{
		repository: repo,
	}
}
