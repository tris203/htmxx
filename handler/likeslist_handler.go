package handler

import (
	"htmxx/dummy"
	"htmxx/service"
	"htmxx/templ"
	"net/http"
)

type LikesListHandler struct {
	userService service.UserService
}

func (h *LikesListHandler) GetLikesList(w http.ResponseWriter, r *http.Request) {
	// handle likes list
	user := r.PathValue("author")
	if user == "" {
		user = h.userService.GetCurrentUser(r)
	}
	currentUser := h.userService.GetCurrentUser(r)
	likes, err := dummy.GetLikedTweets(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	likesListComponent := templ.LikesList(likes, user, currentUser)
	templ.Layout(likesListComponent, "Likes", false).Render(r.Context(), w)
}
