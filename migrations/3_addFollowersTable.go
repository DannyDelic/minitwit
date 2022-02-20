package main

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table follower...")
		_, err := db.Exec(`CREATE TABLE follower(
      account_id INT REFERENCES account(account_id) ON DELETE CASCADE,
      follows_id INT REFERENCES account(account_id) ON DELETE CASCADE
    )`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table follower...")
		_, err := db.Exec(`DROP TABLE follower`)
		return err
	})
}
