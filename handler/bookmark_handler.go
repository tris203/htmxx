package handler

import (
	"htmxx/db"
	"htmxx/service"
	"htmxx/templ"
	"htmxx/model"
	"net/http"
	"strconv"
)

type BookmarkHandler struct {
	userService service.UserService
	tweetService service.TweetService
	dbService   service.DBService
}


func shapeDBBookmarkedTweets(tweets []db.GetBookmarkedTweetsRow) []*model.Tweet {
	var shapedTweets []*model.Tweet
	for _, tweet := range tweets {
		shapedTweets = append(shapedTweets, &model.Tweet{
		ID:               int(tweet.Tweet.TweetID),
		Author:           tweet.Tweet.Author,
		Content:          tweet.Tweet.Content,
		Created:          tweet.Tweet.Created.Time,
		LikeCount:        int(tweet.Tweet.LikeCount),
		LikedBySelf:      tweet.Likedbyuser,
		BookmarkedBySelf: true,
		})
	}
return shapedTweets
}

func (h *BookmarkHandler) GetBookmark(w http.ResponseWriter, r *http.Request) {
	user := h.userService.GetCurrentUser(r)


	ctx, queries, dbConn, dberr := h.dbService.Connect()
	if dberr != nil {
		return
	}
	defer dbConn.Close()

	bookmarks, err := queries.GetBookmarkedTweets(ctx, db.GetBookmarkedTweetsParams{Username: user, Username_2: user})
	if err != nil {
		// handle error
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	bookmarksComponent := templ.BookmarkList(shapeDBBookmarkedTweets(bookmarks), user)
	templ.Layout(bookmarksComponent, "Bookmarks", false).Render(r.Context(), w)
}

func (h *BookmarkHandler) AddBookmark(w http.ResponseWriter, r *http.Request) {
	user := h.userService.GetCurrentUser(r)
	tweetID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := h.tweetService.AddBookmark(tweetID, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookmarkComponent := templ.Bookmark(int(tweetID), result)
	bookmarkComponent.Render(r.Context(), w)
}
