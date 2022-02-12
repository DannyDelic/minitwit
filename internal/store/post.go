package store

import "time"

type Post struct {
	Id       int64
	AuthorId int64
	Text     string
	PubDate  time.Time
	Flagged  int32
}
