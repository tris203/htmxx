package model

import "time"

type Tweet struct {
	ID       int
	Author   string
	IsAuthor bool
	Content  string
	Created  time.Time
	// TODO: this should be an int list of user IDs
	LikedBy     []string
	LikedBySelf bool
	BookmarkedBySelf bool
}
