package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table account...")
		_, err := db.Exec(`CREATE TABLE account(
      account_id SERIAL PRIMARY KEY,
      username TEXT NOT NULL UNIQUE,
      pw_hash TEXT NOT NULL,
	  email TEXT not null,
      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    )`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table account...")
		_, err := db.Exec(`DROP TABLE account`)
		return err
	})
}
