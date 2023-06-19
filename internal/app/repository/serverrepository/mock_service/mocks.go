package mock_service

import (
	"errors"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
)

// MockKeeperRepository is a mock of KeeperRepository interface.
type MockKeeperRepository struct {
}

func NewMockKeeperRepository() *MockKeeperRepository {
	return &MockKeeperRepository{}
}

// AddUser mocks base method.
func (m *MockKeeperRepository) AddUser(userID, login, password string) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// CheckUser mocks base method.
func (m *MockKeeperRepository) CheckUser(login, password string) (string, error) {
	if login == "fail" {
		return "", errors.New("test")
	}
	return "", nil
}

// DeleteData mocks base method.
func (m *MockKeeperRepository) DeleteData(key, userID string) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// SetBankCardData mocks base method.
func (m *MockKeeperRepository) SetBankCardData(userID string, data models.BankCardData) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// SetBinaryData mocks base method.
func (m *MockKeeperRepository) SetBinaryData(userID string, data models.BinaryData) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// SetLoginPasswordData mocks base method.
func (m *MockKeeperRepository) SetLoginPasswordData(userID string, data models.LoginPasswordData) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// SetTextData mocks base method.
func (m *MockKeeperRepository) SetTextData(userID string, data models.TextData) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// Sync mocks base method.
func (m *MockKeeperRepository) Sync(userID string) ([]any, error) {
	if userID == "fail" {
		return nil, errors.New("test")
	}
	return nil, nil
}

// UpdateBankCardData mocks base method.
func (m *MockKeeperRepository) UpdateBankCardData(userID string, data models.BankCardData) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// UpdateBinaryData mocks base method.
func (m *MockKeeperRepository) UpdateBinaryData(userID string, data models.BinaryData) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// UpdateLoginPasswordData mocks base method.
func (m *MockKeeperRepository) UpdateLoginPasswordData(userID string, data models.LoginPasswordData) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}

// UpdateTextData mocks base method.
func (m *MockKeeperRepository) UpdateTextData(userID string, data models.TextData) error {
	if userID == "fail" {
		return errors.New("test")
	}
	return nil
}
