package model

import "time"

type Tweet struct {
	ID       int64
	Author   string
	Content  string
	Created  time.Time
	LikeCount int64
	LikedBySelf bool
	BookmarkedBySelf bool
	ParentID int64
}
