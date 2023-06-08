package service

import (
	"encoding/hex"
	"fmt"

	"github.com/cucumberjaye/GophKeeper/internal/models"
	"github.com/cucumberjaye/GophKeeper/pkg/encryption"
)

func (s *KeeperService) SetOrUpdateLoginPasswordData(userID string, data models.LoginPasswordData) error {
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

	return s.repository.SetOrUpdateLoginPasswordData(userID, data)
}

func (s *KeeperService) SetOrUpdateTextData(userID string, data models.TextData) error {
	var err error
	if data.Data != "" {
		data.Data, err = encryption.Encrypt(data.Data)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}
	return s.repository.SetOrUpdateTextData(userID, data)
}

func (s *KeeperService) SetOrUpdateBinaryData(userID string, data models.BinaryData) error {
	var err error
	if len(data.Data) > 0 {
		data.HexData, err = encryption.EncryptBin(data.Data)
		if err != nil {
			return fmt.Errorf("encryption failed with error: %w", err)
		}
	}
	return s.repository.SetOrUpdateBinaryData(userID, data)
}

func (s *KeeperService) SetOrUpdateBankCardData(userID string, data models.BankCardData) error {
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
	return s.repository.SetOrUpdateBankCardData(userID, data)
}

func (s *KeeperService) GetData(key, userID string) (any, error) {
	data, err := s.repository.GetData(key, userID)
	if err != nil {
		return nil, err
	}

	binData, ok := data.(models.BinaryData)
	if ok {
		binData.Data, err = hex.DecodeString(binData.HexData)
		if err != nil {
			return nil, fmt.Errorf("decode string failed with error: %w", err)
		}
		return binData, nil
	}

	return data, nil
}

func (s *KeeperService) GetDataArray(userID string) ([]string, error) {
	return s.repository.GetDataArray(userID)
}

func (s *KeeperService) DeleteData(key, userID string) error {
	return s.repository.DeleteData(key, userID)
}
