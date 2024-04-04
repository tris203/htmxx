package service

import (
	"htmxx/db"
	"htmxx/model"
)

type TimelineService struct {
	dbService DBService
}

func shapeDBTimeline(tweets []db.GetTimelineRow) []*model.Tweet {
	var shapedTweets []*model.Tweet
	for _, tweet := range tweets {
		shapedTweets = append(shapedTweets, &model.Tweet{
			ID:               int(tweet.Tweet.TweetID),
			Author:           tweet.Tweet.Author,
			Content:          tweet.Tweet.Content,
			Created:          tweet.Tweet.Created.Time,
			LikeCount:        int(tweet.Tweet.LikeCount),
			LikedBySelf:      tweet.Likedbyuser,
			BookmarkedBySelf: tweet.Bookmarkedbyuser,
		})
	}
	return shapedTweets
}

func (s *TimelineService) GetTimeline(minid int64, currentUser string) ([]*model.Tweet, error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return nil, dberr
	}
	defer dbConn.Close()
	tweet, err := queries.GetTimeline(ctx, db.GetTimelineParams{TweetID: minid, Author: "%", Username: currentUser, Username_2: currentUser})

	if err != nil {
		return nil, err
	}
	return shapeDBTimeline(tweet), nil
}

func (s *TimelineService) GetUserTimeline(author string, minid int64, currentUser string) ([]*model.Tweet, error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return nil, dberr
	}
	defer dbConn.Close()

	tweet, err := queries.GetTimeline(ctx, db.GetTimelineParams{TweetID: minid, Author: author, Username: currentUser, Username_2: currentUser})

	if err != nil {
		return nil, err
	}
	return shapeDBTimeline(tweet), nil
}
