package service

import (
	"fmt"
	"htmxx/dummy"
	"htmxx/model"
)

type TweetService struct {
}

func (s *TweetService) CreateTweet(tweet *model.Tweet) (bool, error) {
	success, err := dummy.AddDummyTweet(tweet)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (s *TweetService) GetTweet(id int, userid string) (*model.Tweet, error) {
	var allTweets, err = dummy.GetDummyTweetList(userid)
	if err != nil {
		return nil, err
	}
	for _, tweet := range allTweets {
		if tweet.ID == id {
			return tweet, nil
		}
	}
	return nil, fmt.Errorf("Tweet not found")

}

func (s *TweetService) AddLike(id int, userid string) (likeCount int, likedBySelf bool, err error) {
	success, likedBySelf, err := dummy.AddLike(id, userid)
	if err != nil {
		return 0, false, err
	}
	return success, likedBySelf, nil
}

func (s *TweetService) DeleteTweet(id int, requester string) (tweet model.Tweet, err error) {
	deletedTweet, err := dummy.DeleteTweet(id, requester)
	if err != nil {
		return 
	}
	return deletedTweet, nil
}
