package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func latest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"latest": LATEST,
	})
}
