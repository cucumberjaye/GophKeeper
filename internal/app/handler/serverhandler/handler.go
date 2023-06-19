package serverhandler

import (
	"github.com/cucumberjaye/GophKeeper/internal/app/models"
)

// KeeperService - интерфейс для взаимодействия с логическим слоем.
type KeeperService interface {
	AuthService
	StoreService
}

// AuthService - интерефейс для взаимодействия с логическими методами аутентификации.
type AuthService interface {
	AddUser(login, password string) error
	CreateToken(login, password string) (string, error)
}

// StoreService - интерефейс для взаимодействия с логическими методами храниения и изменения данных.
type StoreService interface {
	SetLoginPasswordData(userID string, data models.LoginPasswordData) error
	SetTextData(userID string, data models.TextData) error
	SetBinaryData(userID string, data models.BinaryData) error
	SetBankCardData(userID string, data models.BankCardData) error
	Sync(userID string) ([]any, error)
	DeleteData(key, userID string) error
	UpdateLoginPasswordData(userID string, data models.LoginPasswordData) error
	UpdateTextData(userID string, data models.TextData) error
	UpdateBinaryData(userID string, data models.BinaryData) error
	UpdateBankCardData(userID string, data models.BankCardData) error
}
