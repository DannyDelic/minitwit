package store

import (
	"errors"
	"time"
)

type Account struct {
	tableName  struct{} `pg:"account,alias:acc"`
	AccountId  int64
	Username   string
	Email      string
	PwHash     string `json:"pw_hash"`
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func AddAccount(account *Account) error {
	_, err := db.Model(account).Returning("*").Insert()
	if err != nil {
		return err
	}
	return nil
}

func Authenticate(username, pwHash string) (*Account, error) {
	account := new(Account)
	if err := db.Model(account).Where(
		"username = ?", username).Select(); err != nil {
		return nil, err
	}
	if pwHash != account.PwHash {
		return nil, errors.New("Password not valid.")
	}
	return account, nil
}
