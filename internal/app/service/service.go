package service

import "github.com/cucumberjaye/GophKeeper/internal/app/models"

type KeeperRepository interface {
	AuthRepository
	StoreRepository
}

type AuthRepository interface {
	AddUser(userID, login, password string) error
	CheckUser(login, password string) (string, error)
}

type StoreRepository interface {
	SetLoginPasswordData(userID string, data models.LoginPasswordData) error
	SetTextData(userID string, data models.TextData) error
	SetBinaryData(userID string, data models.BinaryData) error
	SetBankCardData(userID string, data models.BankCardData) error
	//GetData(key, userID string) (any, error)
	GetDataArray(userID string) ([]any, error)
	DeleteData(key, userID string) error
	UpdateLoginPasswordData(userID string, data models.LoginPasswordData) error
	UpdateTextData(userID string, data models.TextData) error
	UpdateBinaryData(userID string, data models.BinaryData) error
	UpdateBankCardData(userID string, data models.BankCardData) error
}

type KeeperService struct {
	repository KeeperRepository
}

func New(repo KeeperRepository) *KeeperService {
	return &KeeperService{
		repository: repo,
	}
}
