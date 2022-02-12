package server

import (
	"minitwit/internal/database"
	"minitwit/internal/store"
)

func Start() {
	store.SetDBConnection(database.NewDBOptions())

	router := setRouter()

	// Start listening and serving requests
	router.Run(":3000")
}
