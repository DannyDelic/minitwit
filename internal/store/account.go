package store

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Account struct {
	tableName      struct{} `pg:"account,alias:acc"`
	AccountId      int64
	Username       string
	Email          string
	Password       string `pg:"-"`
	HashedPassword []byte `json:"-"`
	Salt           []byte `json:"-"`
	CreatedAt      time.Time
	ModifiedAt     time.Time
}

func AddAccount(account *Account) error {
	salt, err := GenerateSalt()
	if err != nil {
		return err
	}
	toHash := append([]byte(account.Password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	account.Salt = salt
	account.HashedPassword = hashedPassword

	_, err = db.Model(account).Returning("*").Insert()
	if err != nil {
		return err
	}
	return err
}

func FetchAccount(id int) (*Account, error) {
	account := new(Account)
	account.AccountId = int64(id)
	err := db.Model(account).Returning("*").WherePK().Select()
	if err != nil {
		log.Println("Error fetching account")
		return nil, err
	}
	return account, nil
}

func Authenticate(username, password string) (*Account, error) {
	account := new(Account)
	if err := db.Model(account).Where(
		"username = ?", username).Select(); err != nil {
		return nil, err
	}
	salted := append([]byte(password), account.Salt...)
	if err := bcrypt.CompareHashAndPassword(account.HashedPassword, salted); err != nil {
		return nil, err
	}
	return account, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
