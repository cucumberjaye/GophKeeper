package repository

import (
	"context"
	"fmt"

	"github.com/cucumberjaye/GophKeeper/internal/models"
	"github.com/jackc/pgx/v4"
)

func (r *KeeperRepository) SetOrUpdateLoginPasswordData(userID string, data models.LoginPasswordData) error {
	insertTag, err := r.db.Exec(context.Background(), `INSERT INTO login_password (description, login, password, user_id) values($1, $2, $3, $4) 
		ON CONFLICT (description, user_id) DO UPDATE SET login=$2, password=$3`,
		data.Description, data.Login, data.Password, userID)
	if err != nil {
		return fmt.Errorf("insert in login_password table failed with error: %w", err)
	}

	if insertTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected: %w", pgx.ErrNoRows)
	}

	return nil
}

func (r *KeeperRepository) SetOrUpdateTextData(userID string, data models.TextData) error {
	insertTag, err := r.db.Exec(context.Background(), `INSERT INTO text_data (description, data, user_id) values($1, $2, $3) 
	ON CONFLICT (description, user_id) DO UPDATE SET data=$2`,
		data.Description, data.Data, userID)
	if err != nil {
		return fmt.Errorf("insert in text_data table failed with error: %w", err)
	}

	if insertTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected: %w", pgx.ErrNoRows)
	}

	return nil
}

func (r *KeeperRepository) SetOrUpdateBinaryData(userID string, data models.BinaryData) error {
	insertTag, err := r.db.Exec(context.Background(), `INSERT INTO binary_data (description, data, user_id) values($1, $2, $3) 
	ON CONFLICT (description, user_id) DO UPDATE SET data=$2`,
		data.Description, data.HexData, userID)
	if err != nil {
		return fmt.Errorf("insert in binary_data table failed with error: %w", err)
	}

	if insertTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected: %w", pgx.ErrNoRows)
	}

	return nil
}

func (r *KeeperRepository) SetOrUpdateBankCardData(userID string, data models.BankCardData) error {
	insertTag, err := r.db.Exec(context.Background(), `INSERT INTO backcard_data (description, number, valid_thru, cvv, user_id) values($1, $2, $3, $4, $5) 
	ON CONFLICT (description, user_id) DO UPDATE SET number=$2, valid_thru=$3, cvv=$4`,
		data.Description, data.Number, data.ValidThru, data.CVV, userID)
	if err != nil {
		return fmt.Errorf("insert in backcard_data table failed with error: %w", err)
	}

	if insertTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected: %w", pgx.ErrNoRows)
	}

	return nil
}

func (r *KeeperRepository) GetData(key, userID string) (any, error) {
	var tableName string
	err := r.db.QueryRow(context.Background(), `SELECT tablename FROM login_password WHERE user_id=$1 and description=$2 UNION
	SELECT tablename FROM text_data WHERE user_id=$1 and description=$2 UNION
	SELECT tablename FROM binary_data WHERE user_id=$1 and description=$2 UNION
	SELECT tablename FROM backcard_data WHERE user_id=$1 and description=$2`, userID, key).Scan(&tableName)
	if err != nil {
		return nil, fmt.Errorf("select data failed with error: %w", err)
	}

	switch tableName {
	case "login_password":
		var res models.LoginPasswordData
		err := r.db.QueryRow(context.Background(), "SELECT description, login, password FROM login_password WHERE description=$1 and user_id=$2", key, userID).
			Scan(&res.Description, &res.Login, &res.Password)
		if err != nil {
			return nil, fmt.Errorf("select data from login_password table failed with error: %w", err)
		}
		return res, nil
	case "text_data":
		var res models.TextData
		err = r.db.QueryRow(context.Background(), "SELECT description, data FROM text_data WHERE description=$1 and user_id=$2", key, userID).
			Scan(&res.Description, &res.Data)
		if err != nil {
			return nil, fmt.Errorf("select data from text_data table failed with error: %w", err)
		}
		return res, nil
	case "binary_data":
		var res models.BinaryData
		err = r.db.QueryRow(context.Background(), "SELECT description, data FROM binary_data WHERE description=$1 and user_id=$2", key, userID).
			Scan(&res.Description, &res.HexData)
		if err != nil {
			return nil, fmt.Errorf("select data from binary_data table failed with error: %w", err)
		}
		return res, nil
	case "backcard_data":
		var res models.BankCardData
		err = r.db.QueryRow(context.Background(), "SELECT description, number, valid_thru, cvv FROM backcard_data WHERE description=$1 and user_id=$2", key, userID).
			Scan(&res.Description, &res.Number, &res.ValidThru, &res.CVV)
		if err != nil {
			return nil, fmt.Errorf("select data from backcard_data table failed with error: %w", err)
		}
		return res, nil
	}

	return nil, ErrDataNotFound
}

func (r *KeeperRepository) GetDataArray(userID string) ([]string, error) {
	rows, err := r.db.Query(context.Background(), `SELECT description FROM login_password WHERE user_id=$1 UNION
	SELECT description FROM text_data WHERE user_id=$1 UNION
	SELECT description FROM binary_data WHERE user_id=$1 UNION
	SELECT description FROM backcard_data WHERE user_id=$1`, userID)
	if err != nil {
		return nil, fmt.Errorf("select data array failed with error: %w", err)
	}
	defer rows.Close()

	result := []string{}
	for rows.Next() {
		var tmp string
		rows.Scan(&tmp)
		result = append(result, tmp)
	}

	return result, rows.Err()
}

func (r *KeeperRepository) DeleteData(key, userID string) error {
	var tableName string
	err := r.db.QueryRow(context.Background(), `SELECT tablename FROM login_password WHERE user_id=$1 and description=$2 UNION
	SELECT tablename FROM text_data WHERE user_id=$1 and description=$2 UNION
	SELECT tablename FROM binary_data WHERE user_id=$1 and description=$2 UNION
	SELECT tablename FROM backcard_data WHERE user_id=$1 and description=$2`, userID, key).Scan(&tableName)
	if err != nil {
		return fmt.Errorf("select data failed with error: %w", err)
	}

	_, err = r.db.Exec(context.Background(), fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 and description=$2", tableName), userID, key)
	if err != nil {
		return fmt.Errorf("delete data failed with error: %w", err)
	}

	return nil
}
