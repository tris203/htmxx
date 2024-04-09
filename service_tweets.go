package main

import (
	"context"
	"database/sql"
	"fmt"
	"htmxx/db"
	"htmxx/model"
)

func shapeDBTweet(tweet []db.GetTweetRow) []*model.Tweet {
	var shapeTweet []*model.Tweet

	for _, t := range tweet {
		shapeTweet = append(shapeTweet, &model.Tweet{
			ID:               t.Tweet.TweetID,
			Author:           t.Tweet.Author,
			Content:          t.Tweet.Content,
			Created:          t.Tweet.Created.Time,
			LikeCount:        t.Tweet.LikeCount,
			LikedBySelf:      t.Likedbyuser,
			BookmarkedBySelf: t.Bookmarkedbyuser,
		})
	}
	return shapeTweet
}

func (s *application) CreateTweetData(tweet *model.Tweet, ctx context.Context) (newid int64, error error) {

	newid, err := s.query.AddTweet(ctx, db.AddTweetParams{Author: tweet.Author, Content: tweet.Content})
	if err != nil {
		return 0, err
	}
	return newid, nil
}

func (s *application) GetTweetData(id int64, ctx context.Context) ([]*model.Tweet, error) {

	userid := ctx.Value("user").(string)

	tweet, err := s.query.GetTweet(ctx, db.GetTweetParams{TweetID: id, Username: userid, Username_2: userid, ParentTweetID: sql.NullInt64{Int64: id, Valid: true}})

	fmt.Println(tweet)
	fmt.Println(len(tweet))

	if err != nil {
		return nil, err
	}

	return shapeDBTweet(tweet), nil
}

func (s *application) AddLikeData(id int64, ctx context.Context) (likeCount int64, likedBySelf bool, err error) {
	userid := ctx.Value("user").(string)

	likeErr := s.query.LikeTweet(ctx, db.LikeTweetParams{TweetID: id, Username: userid})
	if likeErr != nil {
		return 0, false, fmt.Errorf("You already liked this tweet")
	}
	newLikeCount, err := s.query.GetLikeCount(ctx, id)
	return newLikeCount, true, err
}

func (s *application) RemoveLikeData(id int64, ctx context.Context) (likeCount int64, likedBySelf bool, err error) {
	userid := ctx.Value("user").(string)

	likeErr := s.query.UnlikeTweet(ctx, db.UnlikeTweetParams{TweetID: id, Username: userid})
	if likeErr != nil {
		return 0, false, likeErr
	}
	newLikeCount, err := s.query.GetLikeCount(ctx, id)
	return newLikeCount, false, err
}

func (s *application) DeleteTweetData(id int64, ctx context.Context) (deletedid int64, err error) {
	requester := ctx.Value("user").(string)

	tweet, err := s.query.DeleteTweet(ctx, db.DeleteTweetParams{TweetID: id, Author: requester})
	if err != nil {
		return 0, fmt.Errorf("You are not authorized to delete this tweet")
	}
	if tweet.TweetID == 0 {
		return 0, fmt.Errorf("You are not authorized to delete this tweet")
	}
	return tweet.TweetID, nil
}

func (s *application) AddBookmarkData(id int64, ctx context.Context) (bookmarked bool, err error) {
	userid := ctx.Value("user").(string)

	bookmarkErr := s.query.BookmarkTweet(ctx, db.BookmarkTweetParams{TweetID: id, Username: userid})
	if bookmarkErr != nil {
		// check if the error is a duplicate bookmark error
		return false, bookmarkErr
	}

	return true, nil
}

func (s *application) RemoveBookmarkData(id int64, ctx context.Context) (bookmarked bool, err error) {
	userid := ctx.Value("user").(string)

	bookmarkErr := s.query.UnbookmarkTweet(ctx, db.UnbookmarkTweetParams{TweetID: id, Username: userid})
	if bookmarkErr != nil {
		return true, bookmarkErr
	}
	return false, nil
}

func (s *application) ReplyTweetData(id int64, content string, ctx context.Context) (newid int64, error error) {
	userid := ctx.Value("user").(string)

	newid, err := s.query.ReplyTweet(ctx, db.ReplyTweetParams{Author: userid, Content: content, ParentTweetID: sql.NullInt64{Int64: id, Valid: true}})

	if err != nil {
		return 0, err
	}
	return newid, nil
}
