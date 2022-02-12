package store

import (
	"errors"
	"time"
)

type BingBong struct {
	UserId     int64
	Username   string
	Email      string
	PwHash     string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func AddAccount(account *BingBong) error {
	_, err := db.Model(account).Returning("*").Insert()
	if err != nil {
		return err
	}
	return nil
}

func Authenticate(username, pwHash string) (*BingBong, error) {
	account := new(BingBong)
	if err := db.Model(account).Where(
		"username = ?", username).Select(); err != nil {
		return nil, err
	}
	if pwHash != account.PwHash {
		return nil, errors.New("Password not valid.")
	}
	return account, nil
}
