package store

import (
	"log"
	"time"
)

type Post struct {
	tableName  struct{}  `pg:"post,alias:tbl"`
	PostID     int64     `json:"post_id"`
	AccountID  int64     `json:"-"`
	Poster     string    `json:"poster"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func AddPost(account *Account, post *Post) error {
	post.AccountID = account.AccountID
	_, err := db.Model(post).Returning("*").Insert()
	if err != nil {
		log.Println("Error inserting new post")
	}
	return err
}

func FetchUserPosts(posts *[]Post, id int64) error {
	_, err := db.Query(posts, `select poster, post.content, post.created_at from post, account where post.account_id = account.account_id and account.account_id = ?`, id, id)
	if err != nil {
		log.Println("Error fetching account's posts")
	}
	return err
}

func FetchAllPosts(posts *[]Post) error {
	err := db.Model(posts).Order("tbl.created_at DESC").Select()
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
