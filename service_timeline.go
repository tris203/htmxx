package main

import (
	"context"
	"htmxx/db"
	"htmxx/model"
)

func shapeDBTimeline(tweets []db.GetTimelineRow) []*model.Tweet {
	var shapedTweets []*model.Tweet
	for _, tweet := range tweets {
		shapedTweets = append(shapedTweets, &model.Tweet{
			ID:               tweet.Tweet.TweetID,
			Author:           tweet.Tweet.Author,
			Content:          tweet.Tweet.Content,
			Created:          tweet.Tweet.Created.Time,
			LikeCount:        tweet.Tweet.LikeCount,
			LikedBySelf:      tweet.Likedbyuser,
			BookmarkedBySelf: tweet.Bookmarkedbyuser,
		})
	}
	return shapedTweets
}

func shapeDBAllTimeline(tweets []db.GetAllTimelineRow) []*model.Tweet {
	var shapedTweets []*model.Tweet
	for _, tweet := range tweets {
		shapedTweets = append(shapedTweets, &model.Tweet{
			ID:               tweet.Tweet.TweetID,
			Author:           tweet.Tweet.Author,
			Content:          tweet.Tweet.Content,
			Created:          tweet.Tweet.Created.Time,
			LikeCount:        tweet.Tweet.LikeCount,
			LikedBySelf:      tweet.Likedbyuser,
			BookmarkedBySelf: tweet.Bookmarkedbyuser,
		})
	}
	return shapedTweets
}

func (s *application) GetTimelineData(minid int64, ctx context.Context) ([]*model.Tweet, error) {
	currentUser := ctx.Value("user").(string)
	tweet, err := s.query.GetAllTimeline(ctx, db.GetAllTimelineParams{TweetID: minid, Username: currentUser, Username_2: currentUser})

	if err != nil {
		return nil, err
	}
	return shapeDBAllTimeline(tweet), nil
}

func (s *application) GetUserTimelineData(author string, minid int64, ctx context.Context) ([]*model.Tweet, error) {
	currentUser := ctx.Value("user").(string)

	tweet, err := s.query.GetTimeline(ctx, db.GetTimelineParams{TweetID: minid, Author: author, Username: currentUser, Username_2: currentUser})

	if err != nil {
		return nil, err
	}
	return shapeDBTimeline(tweet), nil
}
