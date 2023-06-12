package clientrepository

import (
	"context"
	"fmt"
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

func (r *ClientStorage) SetLoginPasswordsData(data models.LoginPasswordData, userID string) error {
	return r.rdb.HSet(context.Background(), fmt.Sprintf("%s:%s:%s", loginPasswordData, data.Description, userID), &data).Err()
}

func (r *ClientStorage) SetTextData(data models.TextData, userID string) error {
	return r.rdb.HSet(context.Background(), fmt.Sprintf("%s:%s:%s", textData, data.Description, userID), &data).Err()
}

func (r *ClientStorage) SetBinaryData(data models.BinaryData, userID string) error {
	return r.rdb.HSet(context.Background(), fmt.Sprintf("%s:%s:%s", binaryData, data.Description, userID), &data).Err()
}

func (r *ClientStorage) SetBankCardData(data models.BankCardData, userID string) error {
	return r.rdb.HMSet(context.Background(), fmt.Sprintf("%s:%s:%s", bankCardData, data.Description, userID), &data).Err()
}

func (r *ClientStorage) GetDataArray(userID string) ([]any, error) {
	keys, err := r.rdb.Keys(context.Background(), fmt.Sprintf("*%s", userID)).Result()
	if err != nil {
		return nil, fmt.Errorf("get keys failed with error: %w", err)
	}

	result := make([]any, len(keys))
	for i := range keys {
		switch keys[i][:2] {
		case loginPasswordData:
			var tmp models.LoginPasswordData
			_ = r.rdb.HGetAll(context.Background(), keys[i]).Scan(&tmp)
			result[i] = tmp
		case textData:
			var tmp models.TextData
			_ = r.rdb.HGetAll(context.Background(), keys[i]).Scan(&tmp)
			result[i] = tmp
		case binaryData:
			var tmp models.BinaryData
			_ = r.rdb.HGetAll(context.Background(), keys[i]).Scan(&tmp)
			result[i] = tmp
		case bankCardData:
			var tmp models.BankCardData
			_ = r.rdb.HGetAll(context.Background(), keys[i]).Scan(&tmp)
			result[i] = tmp
		}
	}

	return result, nil
}

func (r *ClientStorage) UpdateLoginPasswordsData(data models.LoginPasswordData, userID string) error {
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

func (r *ClientStorage) UpdateTextData(data models.TextData, userID string) error {
	return r.rdb.HSet(context.Background(), data.Description+":"+userID, &data).Err()
}

func (r *ClientStorage) UpdateBinaryData(data models.BinaryData, userID string) error {
	return r.rdb.HSet(context.Background(), data.Description+":"+userID, &data).Err()
}

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

func (r *ClientStorage) DeleteLoginPasswordData(key string, userID string) error {
	return r.rdb.Del(context.Background(), fmt.Sprintf("%s:%s:%s", loginPasswordData, key, userID)).Err()
}

func (r *ClientStorage) DeleteTextData(key string, userID string) error {
	return r.rdb.Del(context.Background(), fmt.Sprintf("%s:%s:%s", textData, key, userID)).Err()
}

func (r *ClientStorage) DeleteBinaryData(key string, userID string) error {
	return r.rdb.Del(context.Background(), fmt.Sprintf("%s:%s:%s", binaryData, key, userID)).Err()
}

func (r *ClientStorage) DeleteBankCardData(key string, userID string) error {
	return r.rdb.Del(context.Background(), fmt.Sprintf("%s:%s:%s", bankCardData, key, userID)).Err()
}

func (r *ClientStorage) SetLastSync(userID string) error {
	return r.rdb.Set(context.Background(), lastModifiedKey+userID, time.Now().Unix(), 0).Err()
}

func (r *ClientStorage) GetLastSync(userID string) (int64, error) {
	return r.rdb.Get(context.Background(), lastModifiedKey+userID).Int64()
}
