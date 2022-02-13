package server

import (
	"minitwit/internal/conf"
	"minitwit/internal/database"
	"minitwit/internal/store"
)

func Start(cfg conf.Config) {
	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
