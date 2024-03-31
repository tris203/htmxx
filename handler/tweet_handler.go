package handler

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"htmxx/model"
	"htmxx/service"
	"htmxx/templ"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type TweetHandler struct {
	tweetService  service.TweetService
	eventsService service.EventsService
}

func (h *TweetHandler) GetTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		// handle error
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	tweet, err := h.tweetService.GetTweet(id, ip)
	if err != nil {
		// handle error
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
	var ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	// envrypt the ip
	author := fmt.Sprintf("%x", md5.Sum([]byte(ip)))

	tweet := &model.Tweet{
		Author:  author,
		Content: content,
	}

	_, err := h.tweetService.CreateTweet(tweet)
	// show delete on first render
	tweet.IsAuthor = true
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	id, err := strconv.Atoi(r.PathValue("id"))
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newLikeCount, likedBySelf, err := h.tweetService.AddLike(id, ip)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	h.eventsService.AddMessage(service.Event{Data: fmt.Sprintf("%d", newLikeCount), EventName: fmt.Sprintf("like-count-%d", id)})
	heartComponent := templ.Heart(id, likedBySelf, true)
	fmt.Fprintf(w, "%d", newLikeCount)
	heartComponent.Render(r.Context(), w)
}

func (h *TweetHandler) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	tweet, err := h.tweetService.DeleteTweet(id, ip)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// handle success
	deleteTweetComponent := templ.DeletedTweet(tweet)
	deleteTweetComponent.Render(r.Context(), w)

}
