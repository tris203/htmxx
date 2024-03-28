package dummy

import (
	"crypto/md5"
	"fmt"
	"htmxx/model"
	"slices"
	"sort"
	"time"
)

var fakeTweetList = []*model.Tweet{
	{
		ID:      1,
		Author:  "InitialAuthor",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec sit amet fermentum lacus. Morbi ut molestie sem. Donec ut fringilla arcu. Mauris eu cursus neque. Suspendisse tincidunt arcu a metus ornare sodales ut et metus. Aliquam vulputate elit quis odio congue cursus. Ut sagittis nulla quis molestie ullamcorper. Nulla ut urna mollis, condimentum nisl eu, malesuada nibh. Nam dapibus velit turpis, et sodales dui rhoncus pulvinar. Curabitur eget pellentesque nulla. Vestibulum sollicitudin sollicitudin mattis. Suspendisse potenti. Aenean sollicitudin justo eros, ac suscipit metus vehicula nec. Donec molestie nec tellus id eleifend. Aliquam erat volutpat.",
		Created: time.Now(),
	},
	{
		ID:      2,
		Author:  "InitialAuthor",
		Content: "Phasellus commodo scelerisque mollis. In non nisl posuere, finibus neque a, volutpat ligula. Nunc consectetur, erat at feugiat rutrum, massa dui consequat ligula, eget tristique felis nisl vel ligula. Vivamus nec urna efficitur, accumsan odio et, eleifend eros. Integer neque velit, finibus quis elit accumsan, sollicitudin hendrerit diam. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Vestibulum pulvinar pretium justo at posuere. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Proin convallis magna in enim condimentum luctus eu quis leo. Sed tellus quam, euismod at nunc sit amet, auctor rhoncus risus. Nunc nec velit eros. Donec feugiat lacinia quam commodo ornare. Donec fringilla libero non massa facilisis, rhoncus commodo ipsum ultricies.",
		Created: time.Now(),
	},
	{
		ID:      3,
		Author:  "InitialAuthor",
		Content: "Just a really short random tweet",
		Created: time.Now(),
	},
	{
		ID:      4,
		Author:  "InitialAuthor",
		Content: "Another short tweet",
		Created: time.Now(),
	},
	{
		ID:      5,
		Author:  "InitialAuthor",
		Content: "Yet another short tweet",
		Created: time.Now(),
	},
	{
		ID:      6,
		Author:  "InitialAuthor",
		Content: "This is a tweet that is longer than the others. It is a bit longer than the others, but not too long. It is just long enough to be longer than the others. I hope you enjoy reading it.",
		Created: time.Now(),
	},
	{
		ID:      7,
		Author:  "InitialAuthor",
		Content: "This is a tweet that is longer than the others. It is a bit longer than the others, but not too long. It is just long enough to be longer than the others. I hope you enjoy reading it.",
		Created: time.Now(),
	},
	{
		ID:      8,
		Author:  "InitialAuthor",
		Content: "This is a tweet that is longer than the others. It is a bit longer than the others, but not too long. It is just long enough to be longer than the others. I hope you enjoy reading it.",
		Created: time.Now(),
	},
}

var tweetsPerPage =  6

var nextTweetID = len(fakeTweetList) + 1

func GetDummyTweetList(requester string) ([]*model.Tweet, error) {
	allTweets := fakeTweetList
	for _, tweet := range allTweets {
		if slices.Contains(tweet.LikedBy, requester) {
			tweet.LikedBySelf = true
		} else {
			tweet.LikedBySelf = false
		}
	}
	for _, tweet := range allTweets {
		if tweet.Author == fmt.Sprintf("%x", md5.Sum([]byte(requester))) {
			tweet.IsAuthor = true
		} else {
			tweet.IsAuthor = false
		}
	}
	return allTweets, nil
}

func AddDummyTweet(tweet *model.Tweet) (bool, error) {
	tweet.ID = nextTweetID
	nextTweetID++
	tweet.Created = time.Now()
	fakeTweetList = append(fakeTweetList, tweet)
	return true, nil
}

func GetTimeline(requester string, max_id int) ([]*model.Tweet, error) {
	if max_id == 0 {
		max_id = nextTweetID
	}
	allTweets, err := GetDummyTweetList(requester)
	if err != nil {
		return nil, err
	}
	sort.Slice(allTweets, func(i, j int) bool {
		return allTweets[i].Created.After(allTweets[j].Created)
	})

	// retyrn tweets with an ID less than max_id
	var timeline []*model.Tweet
	for _, tweet := range allTweets {

		if tweet.ID < max_id {
			timeline = append(timeline, tweet)
		}
		if len(timeline) >= tweetsPerPage {
			break
		}
	}
	return timeline, nil

}

func GetUserTimeline(requester string, author string, max_id int) ([]*model.Tweet, error) {
	if max_id == 0 {
		max_id = nextTweetID
	}
	allTweets, err := GetDummyTweetList(requester)
	if err != nil {
		return nil, err
	}
	var userTweets []*model.Tweet
	for _, tweet := range allTweets {
		if tweet.Author == author {
			userTweets = append(userTweets, tweet)
		}
	}
	sort.Slice(userTweets, func(i, j int) bool {
		return userTweets[i].Created.After(userTweets[j].Created)
	})

	// return tweets with an ID less than max_id
	var timeline []*model.Tweet
	for _, tweet := range userTweets {
		if tweet.ID < max_id {
			timeline = append(timeline, tweet)
		}
		if len(timeline) >= tweetsPerPage {
			break
		}
	}
	return timeline, nil

}

func DeleteTweet(id int, requester string) (tweet model.Tweet, error error) {
	for i, tweet := range fakeTweetList {
		if tweet.ID == id {
			if tweet.Author != fmt.Sprintf("%x", md5.Sum([]byte(requester))) {
				return *tweet, fmt.Errorf("User %s is not the author of tweet %d", requester, id)
			}
			fakeTweetList = append(fakeTweetList[:i], fakeTweetList[i+1:]...)
			return *tweet, nil
		}
	}
	// TODO: Return error
	return model.Tweet{}, nil
}

func AddLike(id int, userid string) (likeCount int, LikedBySelf bool, error error) {
	fmt.Println(fmt.Sprintf("Adding like to tweet %d by user %s", id, userid))
	for _, tweet := range fakeTweetList {
		if tweet.ID == id {
			if slices.Contains(tweet.LikedBy, userid) {
				userIndex := slices.Index(tweet.LikedBy, userid)
				if userIndex != -1 {
					tweet.LikedBy = append(tweet.LikedBy[:userIndex], tweet.LikedBy[userIndex+1:]...)
				}
				return len(tweet.LikedBy), false, nil
			}
			// TODO: Add user ID to likedBy
			tweet.LikedBy = append(tweet.LikedBy, userid)
			return len(tweet.LikedBy), true, nil
		}
	}
	return 0, false, nil
}
