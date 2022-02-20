package store

import (
	"context"
	"crypto/rand"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Account struct {
	tableName      struct{}  `pg:"account,alias:acc"`
	AccountID      int64     `json:"account_id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `pg:"-" json:"pwd"`
	HashedPassword []byte    `json:"-"`
	Salt           []byte    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	ModifiedAt     time.Time `json:"modified_at"`
	Posts          []*Post   `json:"-" pg:"-"`
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
	account.AccountID = int64(id)
	err := db.Model(account).Returning("*").WherePK().Select()
	if err != nil {
		log.Println("Error fetching account")
		return nil, err
	}
	return account, nil
}

func FetchAccountFromName(name string) (*Account, error) {
	account := new(Account)
	_, err := db.QueryOne(account, `select * from account where account.username = ?`, name)
	if err != nil {
		log.Println("Error fetching account")
		return nil, err
	}
	return account, nil
}
func FetchAccountIdFromName(name string) (int64, error) {
	account, err := FetchAccountFromName(name)
	return account.AccountID, err
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

var _ pg.AfterSelectHook = (*Account)(nil)

func (account *Account) AfterSelect(ctx context.Context) error {
	if account.Posts == nil {
		account.Posts = []*Post{}
	}
	return nil
}
