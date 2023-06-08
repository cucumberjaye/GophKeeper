package handler

import (
	"github.com/cucumberjaye/GophKeeper/internal/models"
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
	SetOrUpdateLoginPasswordData(userID string, data models.LoginPasswordData) error
	SetOrUpdateTextData(userID string, data models.TextData) error
	SetOrUpdateBinaryData(userID string, data models.BinaryData) error
	SetOrUpdateBankCardData(userID string, data models.BankCardData) error
	GetData(key, userID string) (any, error)
	GetDataArray(userID string) ([]string, error)
	DeleteData(key, userID string) error
}
