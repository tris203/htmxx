package model

import "time"

type Tweet struct {
	// TODO: change these to int64
	ID       int64
	Author   string
	Content  string
	Created  time.Time
	LikeCount int64
	LikedBySelf bool
	BookmarkedBySelf bool
}
