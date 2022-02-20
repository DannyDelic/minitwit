package server

import (
	"minitwit/internal/conf"
	"minitwit/internal/database"
	"minitwit/internal/store"
)

var IS_SIM = false

func Start(cfg conf.Config) {
	jwtSetup(cfg)

	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
