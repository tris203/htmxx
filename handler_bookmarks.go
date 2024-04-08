package main

import (
	"htmxx/db"
	"htmxx/templ"
	"htmxx/model"
	"net/http"
	"strconv"
)


func shapeDBBookmarkedTweets(tweets []db.GetBookmarkedTweetsRow) []*model.Tweet {
	var shapedTweets []*model.Tweet
	for _, tweet := range tweets {
		shapedTweets = append(shapedTweets, &model.Tweet{
		ID:               tweet.Tweet.TweetID,
		Author:           tweet.Tweet.Author,
		Content:          tweet.Tweet.Content,
		Created:          tweet.Tweet.Created.Time,
		LikeCount:        tweet.Tweet.LikeCount,
		LikedBySelf:      tweet.Likedbyuser,
		BookmarkedBySelf: true,
		})
	}
return shapedTweets
}

func (h *application) GetBookmark(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)

	bookmarks, err := h.query.GetBookmarkedTweets(r.Context(), db.GetBookmarkedTweetsParams{Username: user, Username_2: user})
	if err != nil {
		// handle error
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	bookmarksComponent := templ.BookmarkList(shapeDBBookmarkedTweets(bookmarks), user)
	templ.Layout(bookmarksComponent, "Bookmarks", false).Render(r.Context(), w)
}

func (h *application) AddBookmark(w http.ResponseWriter, r *http.Request) {
	tweetID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := h.AddBookmarkData(tweetID, r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookmarkComponent := templ.Bookmark(tweetID, result)
	bookmarkComponent.Render(r.Context(), w)
}

func (h *application) RemoveBookmark(w http.ResponseWriter, r *http.Request) {
	tweetID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := h.RemoveBookmarkData(tweetID, r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookmarkComponent := templ.Bookmark(tweetID, result)
	bookmarkComponent.Render(r.Context(), w)
}
