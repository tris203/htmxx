package handler

import (
	"htmxx/db"
	"htmxx/model"
	"htmxx/service"
	"htmxx/templ"
	"net/http"
)

type LikesListHandler struct {
	dbService   service.DBService
}

func shapeDBLikedTweets(tweets []db.GetLikedTweetsRow) []*model.Tweet {
	var shapedTweets []*model.Tweet
	for _, tweet := range tweets {
		shapedTweets = append(shapedTweets, &model.Tweet{
		ID:               tweet.Tweet.TweetID,
		Author:           tweet.Tweet.Author,
		Content:          tweet.Tweet.Content,
		Created:          tweet.Tweet.Created.Time,
		LikeCount:        tweet.Tweet.LikeCount,
		LikedBySelf:      true,
		BookmarkedBySelf: tweet.Bookmarkedbyuser,
		})
	}
return shapedTweets
}

func (h *LikesListHandler) GetLikesList(w http.ResponseWriter, r *http.Request) {
	// handle likes list
	user := r.PathValue("author")

	currentUser := r.Context().Value("user").(string)

	queries, dbConn, dberr := h.dbService.Connect()
	if dberr != nil {
		return
	}
	defer dbConn.Close()

	likes, err := queries.GetLikedTweets(r.Context(), db.GetLikedTweetsParams{Username: user, Username_2: currentUser})
	if err != nil {
		// handle error
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	likesListComponent := templ.LikesList(shapeDBLikedTweets(likes), user, currentUser)
	templ.Layout(likesListComponent, "Likes", false).Render(r.Context(), w)
}
