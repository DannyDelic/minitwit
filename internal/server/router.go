package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var LATEST int

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	// Create API route group
	api := router.Group("/api")
	{
		api.POST("/register", signUp)
		api.POST("/login", signIn)
		api.GET("/msgs", allPosts)
		api.GET("/msgs/:username", accountPosts)
	}

	authorized := api.Group("/")
	authorized.Use(authorization)
	{
		authorized.GET("/latest", latest)
		authorized.GET("/mymsgs", indexPosts) // does the same as accountPosts, but uses the context stored account
		authorized.POST("/msgs/:username", createPost)
		authorized.POST("/fllws/:username", followOrUnfollow)
		authorized.GET("/timeline/:username", createPost)
	}

	router.NoRoute(func(ctx *gin.Context) {
		bytes, _ := ctx.GetRawData()
		log.Println(string(bytes))
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
