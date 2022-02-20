package server

import (
	"errors"
	"log"
	"minitwit/internal/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func authorization(ctx *gin.Context) {
	//bytes, _ := ctx.GetRawData()
	//log.Println("Authorization GET RAW DATA: " + string(bytes))
	account := new(store.Account)
	if !IS_SIM {
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

		account, err = store.FetchAccount(accountID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
	} else {
		fromSimulator := ctx.GetHeader("Authorization")
		if fromSimulator != "Basic c2ltdWxhdG9yOnN1cGVyX3NhZmUh" {
			errorMsg := "You are not authorized to use this resource!"
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errorMsg})
			return
		}
		var err error
		username := ctx.Param("username")
		account, err = store.FetchAccountFromName(username)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
	}
	updateLatest(ctx.Query("latest"))
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

func updateLatest(latest string) {
	latestInt, err := strconv.Atoi(latest)
	if err != nil {
	} else {
		LATEST = latestInt
	}
}
