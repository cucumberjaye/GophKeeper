package service

import "github.com/cucumberjaye/GophKeeper/internal/app/models"

// KeeperRepository - интерфейс взаимодействия с репозиторием.
type KeeperRepository interface {
	AuthRepository
	StoreRepository
}

// AuthRepository - интерфейс ваимодейсвия с репозиторием для аутентификации.
type AuthRepository interface {
	AddUser(userID, login, password string) error
	CheckUser(login, password string) (string, error)
}

// StoreRepository - интерфейс ваимодейсвия с репозиторием для хранения и изменеия данных.
type StoreRepository interface {
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

// KeeperService - структура для взаимодействия с репозиторием.
type KeeperService struct {
	repository KeeperRepository
}

// New - инициализирует структуру KeeperService.
func New(repo KeeperRepository) *KeeperService {
	return &KeeperService{
		repository: repo,
	}
}
