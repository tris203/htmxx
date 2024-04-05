package handler

import (
	"bytes"
	"fmt"
	"htmxx/model"
	"htmxx/service"
	"htmxx/templ"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TweetHandler struct {
	tweetService  service.TweetService
	eventsService service.EventsService
	userService   service.UserService
}

func (h *TweetHandler) GetTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		// handle error
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}
	requester := h.userService.GetCurrentUser(r)
	tweet, err := h.tweetService.GetTweet(id, requester)
	if err != nil {
		// handle error
		fmt.Println(err)
		http.Error(w, "Tweet not found", http.StatusNotFound)
		return
	}
	// handle tweet
	tweetComponent := templ.IndivTweet(tweet)
	err = templ.Layout(tweetComponent, "Tweet", false).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TweetHandler) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var content = r.FormValue("content")
	// envrypt the ip
	author := h.userService.GetCurrentUser(r)

	tweet := &model.Tweet{
		Author:  author,
		Content: content,
	}

	newid, err := h.tweetService.CreateTweet(tweet)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tweet.ID = int(newid)
	tweet.Created = time.Now().UTC()
	// handle success
	var stringTmpl bytes.Buffer
	insertTweetComponent := templ.InsertTweet(tweet)
	insertTweetComponent.Render(r.Context(), &stringTmpl)

	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%s", strings.ReplaceAll(stringTmpl.String(), "\n", "")), EventName: "new-tweet"})
	http.Error(w, "", http.StatusCreated)
	confirmationContainer := templ.CreateNewTweetConfirmation()
	confirmationContainer.Render(r.Context(), w)

}

func (h *TweetHandler) AddLike(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	requester := h.userService.GetCurrentUser(r)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newLikeCount, likedBySelf, err := h.tweetService.AddLike(id, requester)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%d", newLikeCount), EventName: fmt.Sprintf("new-like-count-%d", id)})
	heartComponent := templ.Heart(int(id), likedBySelf, int(newLikeCount), true)
	// fmt.Fprintf(w, "%d", newLikeCount)
	heartComponent.Render(r.Context(), w)
}

func (h *TweetHandler) RemoveLike(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	requester := h.userService.GetCurrentUser(r)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newLikeCount, likedBySelf, err := h.tweetService.RemoveLike(id, requester)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%d", newLikeCount), EventName: fmt.Sprintf("new-like-count-%d", id)})
	heartComponent := templ.Heart(int(id), likedBySelf, int(newLikeCount), false)
	heartComponent.Render(r.Context(), w)
}

func (h *TweetHandler) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requester := h.userService.GetCurrentUser(r)
	deletedID, err := h.tweetService.DeleteTweet(id, requester)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	// handle success
	deleteTweetComponent := templ.DeletedTweet(int(deletedID))
	deleteTweetComponent.Render(r.Context(), w)

}
