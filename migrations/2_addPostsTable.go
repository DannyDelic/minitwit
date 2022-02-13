package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table post...")
		_, err := db.Exec(`CREATE TABLE post(
      post_id SERIAL PRIMARY KEY,
	  poster TEXT NOT NULL,
      content TEXT NOT NULL,
      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      account_id INT REFERENCES account ON DELETE CASCADE
    )`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table post...")
		_, err := db.Exec(`DROP TABLE post`)
		return err
	})
}
