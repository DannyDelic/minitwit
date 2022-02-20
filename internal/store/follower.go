package store

import (
	"log"
)

type Follower struct {
	tableName struct{} `pg:"follower,alias:flw"`
	AccountId int64    `json:"account_id"`
	FollowsId int64    `json:"follows_id"`
	Follow    string   `json:"follow" pg:"-"`
	Unfollow  string   `json:"unfollow" pg:"-"`
}

func Follow(follower *Follower) error {
	_, err := db.Model(follower).Returning("*").Insert()
	if err != nil {
		log.Println("Error inserting new follower, error message: " + err.Error())
	}
	return err
}

func Unfollow(follower *Follower) error {
	_, err := db.QueryOne(follower, "DELETE FROM follower where account_id = ? and follows_id = ?", follower.AccountId, follower.FollowsId)
	if err != nil {
		log.Println("Error deleting follower, error message: " + err.Error())
	}
	return err
}
