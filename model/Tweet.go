package model

import "time"

type Tweet struct {
	// TODO: change these to int64
	ID       int
	Author   string
	Content  string
	Created  time.Time
	LikeCount int
	LikedBySelf bool
	BookmarkedBySelf bool
}
