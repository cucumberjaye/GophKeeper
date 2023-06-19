package serverrepository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
)

// SetLoginPasswordData - функция для сохранения данных логина и пароля в Postgres.
func (r *KeeperRepository) SetLoginPasswordData(userID string, data models.LoginPasswordData) error {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("transaction begin failed with error: %w", err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `INSERT INTO users_descriptions (description, user_id) values($1, $2)`,
		data.Description, userID)
	if err != nil {
		return fmt.Errorf("insert in users_descriptions table failed with error: %w", err)
	}

	_, err = tx.Exec(context.Background(), `INSERT INTO login_password (description, login, password, user_id, last_modified) values($1, $2, $3, $4, $5)`,
		data.Description, data.Login, data.Password, userID, data.LastModified)
	if err != nil {
		return fmt.Errorf("insert in login_password table failed with error: %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("tx commit failed with error: %w", err)
	}

	return nil
}

// SetTextDat - функция для сохранения текстовых данных в Postgres.
func (r *KeeperRepository) SetTextData(userID string, data models.TextData) error {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("transaction begin failed with error: %w", err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `INSERT INTO users_descriptions (description, user_id) values($1, $2)`,
		data.Description, userID)
	if err != nil {
		return fmt.Errorf("insert in users_descriptions table failed with error: %w", err)
	}

	_, err = tx.Exec(context.Background(), `INSERT INTO text_data (description, data, user_id, last_modified) values($1, $2, $3, $4)`,
		data.Description, data.Data, userID, data.LastModified)
	if err != nil {
		return fmt.Errorf("insert in text_data table failed with error: %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("tx commit failed with error: %w", err)
	}

	return nil
}

// SetBinaryData - функция для сохранения бинарных данных в Postgres.
func (r *KeeperRepository) SetBinaryData(userID string, data models.BinaryData) error {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("transaction begin failed with error: %w", err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `INSERT INTO users_descriptions (description, user_id) values($1, $2)`,
		data.Description, userID)
	if err != nil {
		return fmt.Errorf("insert in users_descriptions table failed with error: %w", err)
	}

	_, err = r.db.Exec(context.Background(), `INSERT INTO binary_data (description, data, user_id, last_modified) values($1, $2, $3, $4)`,
		data.Description, data.Data, userID, data.LastModified)
	if err != nil {
		return fmt.Errorf("insert in binary_data table failed with error: %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("tx commit failed with error: %w", err)
	}

	return nil
}

// SetBankCardData - функция для сохранения банковских данных в Postgres.
func (r *KeeperRepository) SetBankCardData(userID string, data models.BankCardData) error {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("transaction begin failed with error: %w", err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `INSERT INTO users_descriptions (description, user_id) values($1, $2)`,
		data.Description, userID)
	if err != nil {
		return fmt.Errorf("insert in users_descriptions table failed with error: %w", err)
	}

	_, err = r.db.Exec(context.Background(), `INSERT INTO bankcard_data (description, number, valid_thru, cvv, user_id, last_modified) values($1, $2, $3, $4, $5, $6)`,
		data.Description, data.Number, data.ValidThru, data.CVV, userID, data.LastModified)
	if err != nil {
		return fmt.Errorf("insert in bankcard_data table failed with error: %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("tx commit failed with error: %w", err)
	}

	return nil
}

// Sync - функция для получения всех данных пользователя из Postgres.
func (r *KeeperRepository) Sync(userID string) ([]any, error) {
	var result []any = []any{}

	rows, err := r.db.Query(context.Background(), `SELECT description, login, password FROM login_password WHERE user_id=$1`, userID)
	if err != nil {
		return nil, fmt.Errorf("select data array failed with error: %w", err)
	}

	for rows.Next() {
		var tmp models.LoginPasswordData
		err := rows.Scan(&tmp.Description, &tmp.Login, &tmp.Password)
		if err != nil {
			return nil, fmt.Errorf("scan failed with error: %w", err)
		}
		result = append(result, tmp)
	}
	rows.Close()

	rows, err = r.db.Query(context.Background(), `SELECT description, data FROM text_data WHERE user_id=$1`, userID)
	if err != nil {
		return nil, fmt.Errorf("select data array failed with error: %w", err)
	}

	for rows.Next() {
		var tmp models.TextData
		err = rows.Scan(&tmp.Description, &tmp.Data)
		if err != nil {
			return nil, fmt.Errorf("scan failed with error: %w", err)
		}
		result = append(result, tmp)
	}
	rows.Close()

	rows, err = r.db.Query(context.Background(), `SELECT description, data FROM binary_data WHERE user_id=$1`, userID)
	if err != nil {
		return nil, fmt.Errorf("select data array failed with error: %w", err)
	}

	for rows.Next() {
		var tmp models.BinaryData
		err = rows.Scan(&tmp.Description, &tmp.Data)
		if err != nil {
			return nil, fmt.Errorf("scan failed with error: %w", err)
		}
		result = append(result, tmp)
	}
	rows.Close()

	rows, err = r.db.Query(context.Background(), `SELECT description, number, valid_thru, cvv FROM bankcard_data WHERE user_id=$1`, userID)
	if err != nil {
		return nil, fmt.Errorf("select data array failed with error: %w", err)
	}

	for rows.Next() {
		var tmp models.BankCardData
		err = rows.Scan(&tmp.Description, &tmp.Number, &tmp.ValidThru, &tmp.CVV)
		if err != nil {
			return nil, fmt.Errorf("scan failed with error: %w", err)
		}
		result = append(result, tmp)
	}
	rows.Close()

	return result, nil
}

// DeleteData - функция для удаления данных из Postgres по ключу.
func (r *KeeperRepository) DeleteData(key, userID string) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM users_descriptions WHERE user_id=$1 and description=$2", userID, key)
	if err != nil {
		return fmt.Errorf("delete data failed with error: %w", err)
	}

	return nil
}

// UpdateLoginPasswordData - функция для изменеия данных логина и пароля в Postgres.
func (r *KeeperRepository) UpdateLoginPasswordData(userID string, data models.LoginPasswordData) error {
	set := []string{}
	values := []any{}
	counter := 1

	if data.Login != "" {
		set = append(set, "login=$"+strconv.Itoa(counter))
		values = append(values, data.Login)
		counter++
	}

	if data.Password != "" {
		set = append(set, "password=$"+strconv.Itoa(counter))
		values = append(values, data.Login)
		counter++
	}

	lastModifiedPlace := counter
	set = append(set, "last_modified=$"+strconv.Itoa(counter))
	values = append(values, data.LastModified)
	counter++

	query := fmt.Sprintf(`UPDATE login_password SET %s WHERE description=$%d and user_id=$%d and last_modified<=$%d`, strings.Join(set, ", "), counter, counter+1, lastModifiedPlace)
	values = append(values, data.Description, userID)

	updateTag, err := r.db.Exec(context.Background(), query, values...)
	if err != nil {
		return fmt.Errorf("update login_password table failed with error: %w", err)
	}

	if updateTag.RowsAffected() == 0 {
		return ErrUpdateLate
	}

	return nil
}

// UpdateTextData - функция для изменения текстовых данных в Postgres.
func (r *KeeperRepository) UpdateTextData(userID string, data models.TextData) error {
	set := []string{}
	values := []any{}
	counter := 1

	if data.Data != "" {
		set = append(set, "data=$"+strconv.Itoa(counter))
		values = append(values, data.Data)
		counter++
	}

	lastModifiedPlace := counter
	set = append(set, "last_modified=$"+strconv.Itoa(counter))
	values = append(values, data.LastModified)
	counter++

	query := fmt.Sprintf(`UPDATE text_data SET %s  WHERE description=$%d and user_id=$%d and last_modified<=$%d`, strings.Join(set, ", "), counter, counter+1, lastModifiedPlace)
	values = append(values, data.Description, userID)

	updateTag, err := r.db.Exec(context.Background(), query, values...)
	if err != nil {
		return fmt.Errorf("update text_data table failed with error: %w", err)
	}

	if updateTag.RowsAffected() == 0 {
		return ErrUpdateLate
	}

	return nil
}

// UpdateBinaryData - функция для изменения бинарных данных в Postgres.
func (r *KeeperRepository) UpdateBinaryData(userID string, data models.BinaryData) error {
	set := []string{}
	values := []any{}
	counter := 1

	if len(data.Data) > 0 {
		set = append(set, "data=$"+strconv.Itoa(counter))
		values = append(values, data.Data)
		counter++
	}

	lastModifiedPlace := counter
	set = append(set, "last_modified=$"+strconv.Itoa(counter))
	values = append(values, data.LastModified)
	counter++

	query := fmt.Sprintf(`UPDATE binary_data SET %s WHERE description=$%d and user_id=$%d and last_modified<=$%d`, strings.Join(set, ", "), counter, counter+1, lastModifiedPlace)
	values = append(values, data.Description, userID)

	updateTag, err := r.db.Exec(context.Background(), query, values...)
	if err != nil {
		return fmt.Errorf("update binary_data table failed with error: %w", err)
	}

	if updateTag.RowsAffected() == 0 {
		return ErrUpdateLate
	}

	return nil
}

// UpdateBankCardData - функция для изменения банковских данных в Postgres.
func (r *KeeperRepository) UpdateBankCardData(userID string, data models.BankCardData) error {
	set := []string{}
	values := []any{}
	counter := 1

	if data.Number != "" {
		set = append(set, "number=$"+strconv.Itoa(counter))
		values = append(values, data.Number)
		counter++
	}

	if data.ValidThru != "" {
		set = append(set, "valid_thru=$"+strconv.Itoa(counter))
		values = append(values, data.ValidThru)
		counter++
	}

	if data.CVV != "" {
		set = append(set, "cvv=$"+strconv.Itoa(counter))
		values = append(values, data.CVV)
		counter++
	}

	lastModifiedPlace := counter
	set = append(set, "last_modified=$"+strconv.Itoa(counter))
	values = append(values, data.LastModified)
	counter++

	query := fmt.Sprintf(`UPDATE bankcard_data SET %s WHERE description=$%d and user_id=$%d and last_modified<=$%d`, strings.Join(set, ", "), counter, counter+1, lastModifiedPlace)
	values = append(values, data.Description, userID)

	updateTag, err := r.db.Exec(context.Background(), query, values...)
	if err != nil {
		return fmt.Errorf("update bankcard_data table failed with error: %w", err)
	}

	if updateTag.RowsAffected() == 0 {
		return ErrUpdateLate
	}

	return nil
}
