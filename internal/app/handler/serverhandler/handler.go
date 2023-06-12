package serverhandler

import (
	"github.com/cucumberjaye/GophKeeper/internal/app/models"
)

type KeeperService interface {
	AuthService
	StoreService
}

type AuthService interface {
	AddUser(login, password string) error
	CreateToken(login, password string) (string, error)
}

type StoreService interface {
	SetLoginPasswordData(userID string, data models.LoginPasswordData) error
	SetTextData(userID string, data models.TextData) error
	SetBinaryData(userID string, data models.BinaryData) error
	SetBankCardData(userID string, data models.BankCardData) error
	Sync(last_sync int64, userID string) ([]any, error)
	DeleteData(key, userID string) error
	UpdateLoginPasswordData(userID string, data models.LoginPasswordData) error
	UpdateTextData(userID string, data models.TextData) error
	UpdateBinaryData(userID string, data models.BinaryData) error
	UpdateBankCardData(userID string, data models.BankCardData) error
}
