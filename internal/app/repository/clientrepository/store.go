package clientrepository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
)

const (
	lastModifiedKey   = "last_modified:"
	loginPasswordData = "lp"
	textData          = "td"
	binaryData        = "bd"
	bankCardData      = "cd"
)

// SetData - функция для сохранения данных в Redis.
func (r *ClientStorage) SetData(data any, userID string) error {
	switch tp := data.(type) {
	case models.LoginPasswordData:
		return r.rdb.HSet(context.Background(), fmt.Sprintf("%s:%s:%s", loginPasswordData, tp.Description, userID), &tp).Err()
	case models.TextData:
		return r.rdb.HSet(context.Background(), fmt.Sprintf("%s:%s:%s", textData, tp.Description, userID), &tp).Err()
	case models.BinaryData:
		return r.rdb.HSet(context.Background(), fmt.Sprintf("%s:%s:%s", binaryData, tp.Description, userID), &tp).Err()
	case models.BankCardData:
		return r.rdb.HMSet(context.Background(), fmt.Sprintf("%s:%s:%s", bankCardData, tp.Description, userID), &tp).Err()
	}

	return errors.New("attempt to set unknown type")
}

// GetDataArray - функция для получения данных из Redis.
func (r *ClientStorage) GetDataArray(userID string) ([]any, error) {
	keys, err := r.rdb.Keys(context.Background(), fmt.Sprintf("*%s", userID)).Result()
	if err != nil {
		return nil, fmt.Errorf("get keys failed with error: %w", err)
	}

	result := make([]any, 0)
	for i := range keys {
		switch keys[i][:2] {
		case loginPasswordData:
			var tmp models.LoginPasswordData
			_ = r.rdb.HGetAll(context.Background(), keys[i]).Scan(&tmp)
			result = append(result, tmp)
		case textData:
			var tmp models.TextData
			_ = r.rdb.HGetAll(context.Background(), keys[i]).Scan(&tmp)
			result = append(result, tmp)
		case binaryData:
			var tmp models.BinaryData
			_ = r.rdb.HGetAll(context.Background(), keys[i]).Scan(&tmp)
			result = append(result, tmp)
		case bankCardData:
			var tmp models.BankCardData
			_ = r.rdb.HGetAll(context.Background(), keys[i]).Scan(&tmp)
			result = append(result, tmp)
		}
	}

	return result, nil
}

// UpdateLoginPasswordData - функция для изменеия логина и пароля в Redis.
func (r *ClientStorage) UpdateLoginPasswordData(data models.LoginPasswordData, userID string) error {
	var old models.LoginPasswordData
	var err error
	if data.Login == "" || data.Password == "" {
		err = r.rdb.HMGet(context.Background(), data.Description+":"+userID, "login", "password").Scan(&old)
		if err != nil {
			return fmt.Errorf("get old data failed with error: %w", err)
		}
	}

	if data.Login == "" {
		data.Login = old.Description
	}
	if data.Password == "" {
		data.Password = old.Password
	}

	return r.rdb.HSet(context.Background(), data.Description+":"+userID, &data).Err()
}

// UpdateTextData - функция для изменеия текстовых данных в Redis.
func (r *ClientStorage) UpdateTextData(data models.TextData, userID string) error {
	return r.rdb.HSet(context.Background(), data.Description+":"+userID, &data).Err()
}

// UpdateBinaryData - функция для изменения бинарных данных в Redis.
func (r *ClientStorage) UpdateBinaryData(data models.BinaryData, userID string) error {
	return r.rdb.HSet(context.Background(), data.Description+":"+userID, &data).Err()
}

// UpdateBankCardData - функция для изменения банковских данных в Redis.
func (r *ClientStorage) UpdateBankCardData(data models.BankCardData, userID string) error {
	var old models.BankCardData
	var err error
	if data.Number == "" || data.ValidThru == "" || data.CVV == "" {
		err = r.rdb.HMGet(context.Background(), data.Description+":"+userID, "number", "valid_thru", "cvv").Scan(&old)
		if err != nil {
			return fmt.Errorf("get old data failed with error: %w", err)
		}
	}

	if data.Number == "" {
		data.Number = old.Number
	}
	if data.ValidThru == "" {
		data.ValidThru = old.ValidThru
	}

	if data.CVV == "" {
		data.CVV = old.CVV
	}

	return r.rdb.HSet(context.Background(), data.Description+":"+userID, &data).Err()
}

// DeleteLoginPasswordData - функция для удаления данных логина и пароля из Redis.
func (r *ClientStorage) DeleteLoginPasswordData(key string, userID string) error {
	return r.rdb.Del(context.Background(), fmt.Sprintf("%s:%s:%s", loginPasswordData, key, userID)).Err()
}

// DeleteTextData - функция для удаления текстовых данных из Redis.
func (r *ClientStorage) DeleteTextData(key string, userID string) error {
	return r.rdb.Del(context.Background(), fmt.Sprintf("%s:%s:%s", textData, key, userID)).Err()
}

// DeleteBinaryData - функция для удаления бинарных данных из Redis.
func (r *ClientStorage) DeleteBinaryData(key string, userID string) error {
	return r.rdb.Del(context.Background(), fmt.Sprintf("%s:%s:%s", binaryData, key, userID)).Err()
}

// DeleteBankCardData - функция для удаления банковских данных из Redis.
func (r *ClientStorage) DeleteBankCardData(key string, userID string) error {
	return r.rdb.Del(context.Background(), fmt.Sprintf("%s:%s:%s", bankCardData, key, userID)).Err()
}

// SetLastSync - функция для сохраниения последней синхронизации с сервером.
func (r *ClientStorage) SetLastSync(userID string) error {
	return r.rdb.Set(context.Background(), lastModifiedKey+userID, time.Now().Unix(), 0).Err()
}

// GetLastSync - функция для получения последней синхронизации с сервером.
func (r *ClientStorage) GetLastSync(userID string) (int64, error) {
	return r.rdb.Get(context.Background(), lastModifiedKey+userID).Int64()
}

// GetAllUserKeys - функция для получения всех клучей от данных из Redis.
func (r *ClientStorage) GetAllUserKeys(userID string) (map[string]func(string, string) error, error) {
	keys, err := r.rdb.Keys(context.Background(), fmt.Sprintf("*%s", userID)).Result()
	if err != nil {
		return nil, fmt.Errorf("get keys failed with error: %w", err)
	}

	res := make(map[string]func(string, string) error, len(keys))
	for i := range keys {
		switch keys[i][:2] {
		case loginPasswordData:
			res[strings.Split(keys[i], ":")[1]] = r.DeleteLoginPasswordData
		case textData:
			res[strings.Split(keys[i], ":")[1]] = r.DeleteTextData
		case binaryData:
			res[strings.Split(keys[i], ":")[1]] = r.DeleteBinaryData
		case bankCardData:
			res[strings.Split(keys[i], ":")[1]] = r.DeleteBankCardData
		}
	}

	return res, nil
}
