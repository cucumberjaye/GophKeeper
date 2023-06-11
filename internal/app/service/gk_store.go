package service

import (
	"fmt"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/pkg/encryption"
)

func (s *KeeperService) SetLoginPasswordData(userID string, data models.LoginPasswordData) error {
	var err error

	data.Login, err = encryption.Encrypt(data.Login)
	if err != nil {
		return fmt.Errorf("encryption failed with error: %w", err)
	}

	data.Password, err = encryption.Encrypt(data.Password)
	if err != nil {
		return fmt.Errorf("encryption failed with error: %w", err)
	}

	return s.repository.SetLoginPasswordData(userID, data)
}

func (s *KeeperService) SetTextData(userID string, data models.TextData) error {
	var err error

	data.Data, err = encryption.Encrypt(data.Data)
	if err != nil {
		return fmt.Errorf("encryption failed with error: %w", err)
	}

	return s.repository.SetTextData(userID, data)
}

func (s *KeeperService) SetBinaryData(userID string, data models.BinaryData) error {
	var err error

	data.Data, err = encryption.EncryptBin(data.Data)
	if err != nil {
		return fmt.Errorf("encryption failed with error: %w", err)
	}

	return s.repository.SetBinaryData(userID, data)
}

func (s *KeeperService) SetBankCardData(userID string, data models.BankCardData) error {
	var err error
	data.Number, err = encryption.Encrypt(data.Number)
	if err != nil {
		return fmt.Errorf("encryption failed with error: %w", err)
	}

	data.ValidThru, err = encryption.Encrypt(data.ValidThru)
	if err != nil {
		return fmt.Errorf("encryption failed with error: %w", err)
	}

	data.CVV, err = encryption.Encrypt(data.CVV)
	if err != nil {
		return fmt.Errorf("encryption failed with error: %w", err)
	}

	return s.repository.SetBankCardData(userID, data)
}

/*func (s *KeeperService) GetData(key, userID string) (any, error) {
	return s.repository.GetData(key, userID)
}*/

func (s *KeeperService) GetDataArray(userID string) ([]any, error) {
	return s.repository.GetDataArray(userID)
}

func (s *KeeperService) DeleteData(key, userID string) error {
	return s.repository.DeleteData(key, userID)
}

func (s *KeeperService) UpdateLoginPasswordData(userID string, data models.LoginPasswordData) error {
	var err error
	if data.Login != "" {
		data.Login, err = encryption.Encrypt(data.Login)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}
	if data.Password != "" {
		data.Password, err = encryption.Encrypt(data.Password)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}

	return s.repository.UpdateLoginPasswordData(userID, data)
}

func (s *KeeperService) UpdateTextData(userID string, data models.TextData) error {
	var err error
	if data.Data != "" {
		data.Data, err = encryption.Encrypt(data.Data)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}
	return s.repository.UpdateTextData(userID, data)
}

func (s *KeeperService) UpdateBinaryData(userID string, data models.BinaryData) error {
	var err error
	if len(data.Data) > 0 {
		data.Data, err = encryption.EncryptBin(data.Data)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}
	return s.repository.UpdateBinaryData(userID, data)
}

func (s *KeeperService) UpdateBankCardData(userID string, data models.BankCardData) error {
	var err error
	if data.Number != "" {
		data.Number, err = encryption.Encrypt(data.Number)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}
	if data.ValidThru != "" {
		data.ValidThru, err = encryption.Encrypt(data.ValidThru)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}
	if data.CVV != "" {
		data.CVV, err = encryption.Encrypt(data.CVV)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}
	return s.repository.UpdateBankCardData(userID, data)
}
