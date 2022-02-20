package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"minitwit/internal/store"
	"net/http"
)

func followOrUnfollow(ctx *gin.Context) {
	//ByteBody, _ := ioutil.ReadAll(ctx.Request.Body)
	//ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))
	//log.Println(string(ByteBody))

	follower := new(store.Follower)
	if err := ctx.Bind(follower); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := currentAccount(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	follower.AccountId = account.AccountID

	if len(follower.Follow) != 0 {
		follow(ctx, follower)
	} else {
		unfollow(ctx, follower)
	}
}

func follow(ctx *gin.Context, follower *store.Follower) {
	followsId, err := store.FetchAccountIdFromName(follower.Follow)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	follower.FollowsId = followsId

	err = store.Follow(follower)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Followed successfully.",
		"data": follower,
	})
}

func unfollow(ctx *gin.Context, follower *store.Follower) {
	followsId, err := store.FetchAccountIdFromName(follower.Unfollow)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	follower.FollowsId = followsId

	err = store.Unfollow(follower)
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"You are trying to unfollow someone you are not following": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Unfollowed successfully.",
		"data": follower,
	})
}
