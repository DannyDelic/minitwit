package server

import (
	"minitwit/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createPost(ctx *gin.Context) {
	post := new(store.Post)
	if err := ctx.Bind(post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := currentAccount(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := store.AddPost(account, post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Post created successfully.",
		"data": post,
	})
}

func indexPosts(ctx *gin.Context) {
	user, err := currentAccount(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var posts []store.Post
	if err := store.FetchUserPosts(&posts, user.AccountID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Posts fetched successfully.",
		"data": posts,
	})
}

func accountPosts(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}

	var posts []store.Post
	if err := store.FetchUserPosts(&posts, int64(id)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Posts fetched successfully.",
		"data": posts,
	})
}

func allPosts(ctx *gin.Context) {
	var posts []store.Post
	if err := store.FetchAllPosts(&posts); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Posts fetched successfully.",
		"data": posts,
	})
}
