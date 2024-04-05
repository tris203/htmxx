package service

import (
	"fmt"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"htmxx/db"
	"htmxx/model"
)

type TweetService struct {
	dbService DBService
}

func shapeDBTweet(tweet db.GetTweetRow) *model.Tweet {
	return &model.Tweet{
		ID:               tweet.Tweet.TweetID,
		Author:           tweet.Tweet.Author,
		Content:          tweet.Tweet.Content,
		Created:          tweet.Tweet.Created.Time,
		LikeCount:        tweet.Tweet.LikeCount,
		LikedBySelf:      tweet.Likedbyuser,
		BookmarkedBySelf: tweet.Bookmarkedbyuser,
	}

}

func (s *TweetService) CreateTweet(tweet *model.Tweet) (newid int64, error error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return 0, dberr
	}
	defer dbConn.Close()

	newid, err := queries.AddTweet(ctx, db.AddTweetParams{Author: tweet.Author, Content: tweet.Content})
	if err != nil {
		return 0, err
	}
	return newid, nil
}

func (s *TweetService) GetTweet(id int64, userid string) (*model.Tweet, error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return nil, dberr
	}
	defer dbConn.Close()

	tweet, err := queries.GetTweet(ctx, db.GetTweetParams{TweetID: id, Username: userid, Username_2: userid})

	if err != nil {
		return nil, err
	}

	return shapeDBTweet(tweet), nil
}

func (s *TweetService) AddLike(id int64, userid string) (likeCount int64, likedBySelf bool, err error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return 0, false, dberr
	}
	defer dbConn.Close()

	likeErr := queries.LikeTweet(ctx, db.LikeTweetParams{TweetID: id, Username: userid})
	if likeErr != nil {
		return 0, false, fmt.Errorf("You already liked this tweet")
	}
	newLikeCount, err := queries.GetLikeCount(ctx, id)
	return newLikeCount, true, err
}

func (s *TweetService) RemoveLike(id int64, userid string) (likeCount int64, likedBySelf bool, err error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return 0, false, dberr
	}
	defer dbConn.Close()

	likeErr := queries.UnlikeTweet(ctx, db.UnlikeTweetParams{TweetID: id, Username: userid})
	if likeErr != nil {
		return 0, false, likeErr
	}
	newLikeCount, err := queries.GetLikeCount(ctx, id)
	return newLikeCount, false, err
}

func (s *TweetService) DeleteTweet(id int64, requester string) (deletedid int64, err error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return 0, dberr
	}
	defer dbConn.Close()

	delerr := queries.DeleteTweet(ctx, db.DeleteTweetParams{TweetID: id, Author: requester})
	if delerr != nil {
		return 0, delerr
	}
	return id, nil
}

func (s *TweetService) AddBookmark(id int64, userid string) (bookmarked bool, err error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return false, dberr
	}
	defer dbConn.Close()

	bookmarkErr := queries.BookmarkTweet(ctx, db.BookmarkTweetParams{TweetID: id, Username: userid})
	if bookmarkErr != nil {
		// check if the error is a duplicate bookmark error
		return false, bookmarkErr
	}

	return true, nil
}

func (s *TweetService) RemoveBookmark(id int64, userid string) (bookmarked bool, err error) {
	ctx, queries, dbConn, dberr := s.dbService.Connect()
	if dberr != nil {
		return true, dberr
	}
	defer dbConn.Close()

	bookmarkErr := queries.UnbookmarkTweet(ctx, db.UnbookmarkTweetParams{TweetID: id, Username: userid})
	if bookmarkErr != nil {
		return true, bookmarkErr
	}
	return false, nil
}
