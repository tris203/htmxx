package main

import (
	"htmxx/db"
	"htmxx/model"
	"htmxx/templ"
	"net/http"
)

func shapeDBLikedTweets(tweets []db.GetLikedTweetsRow) []*model.Tweet {
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

func (h *application) GetLikesList(w http.ResponseWriter, r *http.Request) {
	// handle likes list
	currentUser := r.Context().Value("user").(string)
	user := r.PathValue("author")
	if user == "" {
		user = currentUser
	}

	likes, err := h.query.GetLikedTweets(r.Context(), db.GetLikedTweetsParams{Username: currentUser, Username_2: currentUser, Username_3: user})
	if err != nil {
		// handle error
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	likesListComponent := templ.LikesList(shapeDBLikedTweets(likes), user, currentUser)
	templ.Layout(likesListComponent, "Likes", false).Render(r.Context(), w)
}
