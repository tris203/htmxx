package handler

import (
	"htmxx/dummy"
	"htmxx/service"
	"htmxx/templ"
	"net/http"
	"strconv"
)

type BookmarkHandler struct {
	userService service.UserService
}

func (h *BookmarkHandler) GetBookmark(w http.ResponseWriter, r *http.Request) {
	user := h.userService.GetCurrentUser(r)
	bookmarks, err := dummy.GetBookmarkedTweets(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bookmarksComponent := templ.BookmarkList(bookmarks, user)
	templ.Layout(bookmarksComponent, "Bookmarks", false).Render(r.Context(), w)
}

func (h *BookmarkHandler) AddBookmark(w http.ResponseWriter, r *http.Request) {
	user := h.userService.GetCurrentUser(r)
	tweetID, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := dummy.BookmarkTweet(tweetID, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookmarkComponent := templ.Bookmark(tweetID, result)
	bookmarkComponent.Render(r.Context(), w)
}
