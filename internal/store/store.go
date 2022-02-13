package store

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

// Database connector
var db *pg.DB

func SetDBConnection(dbOpts *pg.Options) {
	if dbOpts == nil {
		log.Panicln("DB options can't be nil")
	} else {
		db = pg.Connect(dbOpts)
		db.AddQueryHook(dbLogger{})
	}
}

func GetDBConnection() *pg.DB { return db }

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	bytes, _ := q.FormattedQuery()
	log.Println(string(bytes))
	return nil
}
