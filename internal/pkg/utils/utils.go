package utils

import (
	"fmt"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/pkg/encryption"
)

func PrintLoginPasswordData(data models.LoginPasswordData) error {
	login, err := encryption.Decode(data.Login)
	if err != nil {
		return fmt.Errorf("decode login failed with error: %w", err)
	}
	password, err := encryption.Decode(data.Password)
	if err != nil {
		return fmt.Errorf("decode password failed with error: %w", err)
	}
	fmt.Printf("\nDescription: %s\nLogin: %s\nPassword: %s\n\n", data.Description, login, password)

	return nil
}

func PrintTextData(data models.TextData) error {
	text, err := encryption.Decode(data.Data)
	if err != nil {
		return fmt.Errorf("decode text failed with error: %w", err)
	}

	fmt.Printf("\nDescription: %s\nData: %s\n\n", data.Description, text)

	return nil
}

func PrintBinaryData(data models.BinaryData) error {
	binary, err := encryption.DecodeBin(data.Data)
	if err != nil {
		return fmt.Errorf("decode binary data failed with error: %w", err)
	}

	fmt.Printf("\nDescription: %s\nData: %v\n\n", data.Description, binary)

	return nil
}

func PrintBankCardData(data models.BankCardData) error {
	number, err := encryption.Decode(data.Number)
	if err != nil {
		return fmt.Errorf("decode card number failed with error: %w", err)
	}

	validThru, err := encryption.Decode(data.ValidThru)
	if err != nil {
		return fmt.Errorf("decode card validThru failed with error: %w", err)
	}

	cvv, err := encryption.Decode(data.CVV)
	if err != nil {
		return fmt.Errorf("decode card cvv failed with error: %w", err)
	}

	fmt.Printf("\nDescription: %s\nNumber: %v\nValidThru: %s\nCVV: %s\n\n", data.Description, number, validThru, cvv)

	return nil
}
