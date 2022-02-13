package server

import (
	"minitwit/internal/security"
	"minitwit/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

var HashPassword = security.HashPassword

func signUp(ctx *gin.Context) {
	account := new(store.Account)
	if err := ctx.Bind(account); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := store.AddAccount(account); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":        "Signed up successfully.",
		"jwt":        generateJWT(account),
		"email":      account.Email,
		"account_id": account.AccountID,
		"username":   account.Username,
	})
}

func signIn(ctx *gin.Context) {
	account := new(store.Account)
	if err := ctx.Bind(account); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := store.Authenticate(account.Username, account.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":        "Signed in successfully.",
		"jwt":        generateJWT(account),
		"email":      account.Email,
		"account_id": account.AccountID,
		"username":   account.Username,
	})
}
