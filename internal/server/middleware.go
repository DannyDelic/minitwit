package server

import (
	"errors"
	"log"
	"minitwit/internal/store"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func authorization(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing."})
		return
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is not valid."})
		return
	}
	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing bearer part."})
		return
	}
	accountID, err := verifyJWT(headerParts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	account, err := store.FetchAccount(accountID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.Set("account", account)
	ctx.Next()
}

func currentAccount(ctx *gin.Context) (*store.Account, error) {
	var err error
	_account, exists := ctx.Get("account")
	if !exists {
		err = errors.New("Current context account not set")
		log.Println("Current context account not set")
		return nil, err
	}
	account, ok := _account.(*store.Account)
	if !ok {
		err = errors.New("Context account is not valid type")
		log.Println("Context account is not valid type")
		return nil, err
	}
	return account, nil
}
